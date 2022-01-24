(function () {
    let socket;

    $(() => {
        let url = new URL(window.location.href);
        url.protocol = 'ws';
        url.pathname = url.pathname + '.ws';
        socket = new Socket(url.href);
        setupEvents();
    });

    function setupEvents() {
        $('input').off('click');
        $('input[type="checkbox"]').on('click', function () {
            socket.SaveGroupItem(
                $(this).prop('checked'),
                $(this).data('group'),
                $(this).data('permission'),
            )
        });

        $('input[type="button"]').on('click', function () {
            if ($(this).val() == 'add') {
                socket.AddGroupItem(
                    $('input[type="text"').val(),
                )
            } else if ($(this).val() == 'delete') {
                let name = $(this).parent().parent().find('td:nth-of-type(2)').html();
                if (confirm('Are you Sure you want to delete ' + name)) {
                    socket.DeleteGroupItem(
                        $(this).data("group"),
                    )
                }
            }
        });
    }

    class Socket {
        #ws;
        constructor(url) {
            $('#socketInfo').show()
            this.#ws = new WebSocket(url);
            this.#ws.addEventListener('open', (e) => { this.OpenListener(e); });
            this.#ws.addEventListener('close', (e) => { this.CloseListener(e); });
            this.#ws.addEventListener('error', (e) => { this.ErrorListener(e); });
            this.#ws.addEventListener('message', (e) => { this.MessageListener(e) });

            $(window).bind('beforeunload', (e) => {
                if (this.#ws.readyState == WebSocket.OPEN) {
                    this.#ws.close();
                }
            });
        }


        OpenListener(e) {
            $('#socketInfo').addClass('connected');
        };

        ErrorListener(e) {
            console.error(e);
        }

        CloseListener(e) {
            $('#socketInfo').removeClass('connected');
        };

        SaveGroupItem(checked, groupid, permissionid) {
            this.#ws.send(JSON.stringify({
                command: 'setgroup',
                checked: checked,
                groupid: groupid,
                permissionid: permissionid,
            }));
        }

        AddGroupItem(name) {
            this.#ws.send(JSON.stringify({
                command: 'addgroup',
                name: name,
            }));
        }

        DeleteGroupItem(groupid) {
            this.#ws.send(JSON.stringify({
                command: 'deletegroup',
                groupid: groupid,
            }));
        }

        MessageListener(e) {
            let edata = JSON.parse(e.data);
            let data = edata['data'];
            switch (data['command']) {
                case 'setgroup':
                    break;
                case 'addgroup':
                    if (edata['status'] != 200) {
                        alert(edata['error_message']);
                        break;
                    }
                    $('input[type="text"').val("");
                    let $template = $($('template').html());
                    $template.prop('id', `group${data['groupid']}`)
                    $template.find('td:first-of-type').html(data['groupid'])
                    $template.find('td:nth-of-type(2)').html(data['name']);
                    $template.find('input').data('group', data['groupid']);
                    $template.insertBefore('#zgroup');
                    setupEvents();
                    sortTablebyID();
                    break;
                case 'deletegroup':
                    $('#group' + data['groupid']).remove();
                    sortTablebyID();
                    break;
                default:
                    console.log('ERROR COMMAND NOT KNOWN SKIPPING:', e.data);
                    break;
            }

        };
    }


    function sortTablebyID() {
        var $tbody = $('table tbody');

        $tbody.find('tr').sort(function (a, b) {
            let aid = parseInt($(a).attr('id').replace('group', ''));
            let bid = parseInt($(b).attr('id').replace('group', ''));
            return aid > bid ? 1 : aid < bid ? -1 : 0;
        }).appendTo($tbody);
    }
})();
