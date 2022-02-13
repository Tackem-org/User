(function () {
    $(() => {
        socket.AddReturnAction("user.admin.group.add", GroupAddReturn);
        socket.AddReturnAction("user.admin.group.delete", GroupDeleteReturn);
        SetupEvents();
    });

    function SetupEvents() {
        $('input').off('click');
        $('input[type="checkbox"]').on('click', function () {
            socket.Send({
                command: 'user.admin.group.set',
                checked: $(this).prop('checked'),
                groupid: $(this).data('group'),
                permissionid: $(this).data('permission'),
            });
        });

        $('input[type="button"]').on('click', function () {
            if ($(this).val() == 'add') {
                socket.Send({
                    command: 'user.admin.group.add',
                    name: $('input[type="text"').val(),
                });
            } else if ($(this).val() == 'delete') {
                let name = $(this).parent().parent().find('td:nth-of-type(2)').html();
                if (confirm('Are you Sure you want to delete ' + name)) {
                    socket.Send({
                        command: 'user.admin.group.delete',
                        groupid: $(this).data("group"),
                    });
                }
            }
        });
    }

    function GroupAddReturn(data) {
        if (data['status'] != 200) {
            alert(data['error_message']);
            return;
        }
        $('input[type="text"').val("");
        let $template = $($('template.group').html());
        $template.prop('id', `group${data['data']['groupid']}`)
        $template.find('td:first-of-type').html(data['data']['groupid'])
        $template.find('td:nth-of-type(2)').html(data['data']['name']);
        $template.find('input').data('group', data['data']['groupid']);
        $template.insertBefore('#zgroup');
        SetupEvents();
        SortTablebyID();
    }

    function GroupDeleteReturn(data) {
        $('#group' + data['data']['groupid']).remove();
        SortTablebyID();
    }

    function SortTablebyID() {
        var $tbody = $('table tbody');
        $tbody.find('tr').sort(function (a, b) {
            let aid = parseInt($(a).attr('id').replace('group', ''));
            let bid = parseInt($(b).attr('id').replace('group', ''));
            return aid > bid ? 1 : aid < bid ? -1 : 0;
        }).appendTo($tbody);
    }
})();
