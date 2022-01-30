(function () {
    let socket;

    let timer = 5; //seconds
    let usernameTimeout = null;
    let passwordTimeout = null;

    $(() => {
        let a = window.location.href.split("/");
        a.pop();
        let url = new URL(a.join("/") + "user");
        url.protocol = 'ws';
        url.pathname = url.pathname + '.ws';
        socket = new Socket(url.href);
        setupEvents();
    });

    function setupEvents() {
        $("input").off("click");

        $('#username').on('input', function () {
            if ($(this).val() == "") {
                return
            }
            if (usernameTimeout != null) {
                clearTimeout(usernameTimeout);
                delete (usernameTimeout);
                usernameTimeout = null;
            }
            usernameTimeout = setTimeout(() => { socket.ChangeUsername($(this).val()) }, 1000 * timer);
        });

        $('#acceptusernamerequest').on('click', function () {
            socket.AcceptUsernameChange();
        });

        $('#rejectusernamerequest').on('click', function () {
            socket.RejectUsernameChange();
        });

        $('#password').on('input', function () {
            if ($(this).val() == "") {
                return
            }
            if (passwordTimeout != null) {
                clearTimeout(passwordTimeout);
                delete (passwordTimeout);
                passwordTimeout = null;
            }
            passwordTimeout = setTimeout(() => { socket.ChangePassword($(this).val()) }, 1000 * timer);
        });

        $('[name="disabled"]').on('click', function () {
            socket.ChangeDisabled(
                $(this).prop('checked'),
            )
        });

        $('[name="isadmin"]').on('click', function () {
            socket.ChangeIsAdmin(
                $(this).prop('checked'),
            )
        });

        $('[name="icon"]').on('change', function () {
            var reader = new FileReader();
            reader.onloadend = function () {
                socket.UploadIconBase64(reader.result);
            }
            reader.readAsDataURL(this.files[0]);
        });

        $('[name="clearicon"]').on('click', function () {
            if (confirm('Are you Sure you want to clear the icon')) {
                socket.ClearIcon()
            }
        });

        $('[name="delete"]').on('click', function () {
            if (confirm('Are you Sure you want to delete this user')) {
                socket.DeleteUser()
            }
        });



        $('[name="groups[]"]').on('click', function () {
            socket.ChangeGroup(
                $(this).prop('checked'),
                $(this).data('group'),
            )
        });

        $('[name="permissions[]"]').on('click', function () {
            socket.ChangePermission(
                $(this).prop('checked'),
                $(this).data('permission'),
            )
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

        parseUserID() {
            return parseInt(window.location.href.split("/").pop());
        }

        ChangeUsername(username) {
            this.#ws.send(JSON.stringify({
                command: 'changeusername',
                userid: this.parseUserID(),
                username: username
            }));
        }

        AcceptUsernameChange() {
            this.#ws.send(JSON.stringify({
                command: 'acceptusernamechange',
                userid: this.parseUserID(),
            }));
        }

        RejectUsernameChange() {
            this.#ws.send(JSON.stringify({
                command: 'rejectusernamechange',
                userid: this.parseUserID(),
            }));
        }

        ChangePassword(password) {
            this.#ws.send(JSON.stringify({
                command: 'changepassword',
                userid: this.parseUserID(),
                password: password,
            }));
        }

        ChangeDisabled(checked) {
            this.#ws.send(JSON.stringify({
                command: 'changedisabled',
                userid: this.parseUserID(),
                checked: checked,
            }));
        }

        ChangeIsAdmin(checked) {
            this.#ws.send(JSON.stringify({
                command: 'changeisadmin',
                userid: this.parseUserID(),
                checked: checked,
            }));
        }
        UploadIconBase64(icon) {
            this.#ws.send(JSON.stringify({
                command: 'uploadiconbase64',
                userid: this.parseUserID(),
                icon: icon,
            }));
        }
        ClearIcon() {
            this.#ws.send(JSON.stringify({
                command: 'clearicon',
                userid: this.parseUserID(),
            }));
        }
        DeleteUser() {
            this.#ws.send(JSON.stringify({
                command: 'deleteuser',
                userid: this.parseUserID(),
            }));
        }
        ChangeGroup(checked, group) {
            this.#ws.send(JSON.stringify({
                command: 'changegroup',
                userid: this.parseUserID(),
                checked: checked,
                group: group,
            }));
        }

        ChangePermission(checked, permission) {
            this.#ws.send(JSON.stringify({
                command: 'changepermission',
                userid: this.parseUserID(),
                checked: checked,
                permission: permission,
            }));
        }

        MessageListener(e) {
            let edata = JSON.parse(e.data);
            console.log(edata);
            let data = edata['data'];
            if (edata['status'] != 200) {
                alert(edata['error_message']);
                return;
            }
            switch (data['command']) {
                case 'changeusername':
                    break;
                case 'acceptusernamechange':
                    console.log(data);
                    $('#username').val(data['name']);
                case 'rejectusernamechange':
                    $('#usernamerequest').remove();
                    break;
                case 'changepassword':
                    break;
                case 'changedisabled':
                    break;
                case 'changeisadmin':
                    break;
                case 'uploadiconbase64':
                    $('#usericon').attr('src', data['name']);
                    break
                case 'clearicon':
                    $('#usericon').attr('src', '');
                    $('[name="icon"]').val(null);
                    break
                case 'deleteuser':
                    break;
                case 'changegroup':
                    break;
                case "changepermission":
                    break;
                default:
                    console.log('ERROR COMMAND NOT KNOWN SKIPPING:', e.data);
                    break;
            }

            $('#updatedat').html(data['updatedat']);
        };
    }
})();
