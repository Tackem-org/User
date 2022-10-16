(function () {
    var ready = (callback) => {
        if (document.readyState != "loading") callback();
        else document.addEventListener("DOMContentLoaded", callback);
    };

    ready(() => {
        socket.AddReturnAction("user.admin.group.add", GroupAddReturn);
        socket.AddReturnAction("user.admin.group.delete", GroupDeleteReturn);
        document.querySelectorAll('input[type="checkbox"]').forEach((e) => {
            e.addEventListener('click', ClickCheckbox);
        });

        document.querySelectorAll('input[type="button"]').forEach((e) => {

            e.addEventListener('click', ClickButton);
        });
    });

    function ClickCheckbox() {
        console.log("G", parseInt(this.getAttribute('data-group')), "P", parseInt(this.getAttribute('data-permission')));
        socket.Send({
            command: 'user.admin.group.set',
            checked: this.checked,
            groupid: parseInt(this.getAttribute('data-group')),
            permissionid: parseInt(this.getAttribute('data-permission')),
        });
    }

    function ClickButton() {
        if (this.value == 'add') {
            socket.Send({
                command: 'user.admin.group.add',
                name: document.querySelector('input[type="text"]').value,
            });
        } else if (this.value == 'delete') {
            let name = this.closest('tr').querySelector('td:nth-of-type(2)').innerHTML;
            if (confirm('Are you Sure you want to delete ' + name)) {
                socket.Send({
                    command: 'user.admin.group.delete',
                    groupid: parseInt(this.getAttribute("data-group")),
                });
            }
        }
    }

    function GroupAddReturn(data) {
        if (data['statuscode'] != 200) {
            alert(data['error_message']);
            return;
        }
        document.querySelector('input[type="text"]').value = '';

        let basetemplate = document.querySelector('template.group').content.cloneNode(true);
        let template = basetemplate.querySelector('tr')
        template.setAttribute('id', `group${data['data']['groupid']}`)
        template.querySelector('td:first-of-type').innerHTML = data['data']['groupid'];
        template.querySelector('td:nth-of-type(2)').innerHTML = data['data']['name'];
        template.querySelectorAll('input').forEach((e) => {
            e.setAttribute('data-group', data['data']['groupid']);
        });

        let location = document.getElementById('zgroup')
        location.parentNode.insertBefore(template, location);

        let elem = document.getElementById('group' + data['data']['groupid']);
        console.log("ELEM:", elem);
        elem.querySelectorAll('input[type="checkbox"]').forEach((e) => {
            console.log("CHECKBOXES");
            e.addEventListener('click', ClickCheckbox);
        });

        elem.querySelectorAll('input[type="button"]').forEach((e) => {
            console.log("BUTTONS");
            e.addEventListener('click', ClickButton);
        });
        SortTablebyID();
    }

    function GroupDeleteReturn(data) {
        document.getElementById('group' + data['data']['groupid']).remove();
        SortTablebyID();
    }

    function SortTablebyID() {
        var tbody = document.querySelector('table tbody');
        var rows = [].slice.call(tbody.querySelectorAll("tr"));

        rows.sort(function (a, b) {
            let aid = parseInt(a.getAttribute('id').replace('group', ''));
            let bid = parseInt(b.getAttribute('id').replace('group', ''));
            return aid > bid ? 1 : aid < bid ? -1 : 0;
        });

        rows.forEach(function (v) {
            tbody.appendChild(v);
        });
    }
})();
