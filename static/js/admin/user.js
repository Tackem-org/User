(function () {
    let timer = 5; //seconds
    let usernameTimeout = null;
    let passwordTimeout = null;
    let userID = parseInt(window.location.href.split("/").pop());

    $(() => {
        socket.AddReturnFunction(UpdateTimestamp);
        socket.AddReturnAction("user.acceptusernamechange", AcceptUsernameChange);
        socket.AddReturnAction("user.rejectusernamechange", RejectUsernameChange);
        socket.AddReturnAction("user.admin.edituser.uploadiconbase64", UploadIconBase64);
        socket.AddReturnAction("user.admin.edituser.clearicon", ClearIcon);
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
            usernameTimeout = setTimeout(() => {
                socket.Send({
                    command: 'user.admin.edituser.changeusername',
                    userid: userID,
                    username: $(this).val()
                });
            }, 1000 * timer);
        });

        $('#acceptusernamerequest').on('click', function () {
            socket.Send({
                command: 'user.acceptusernamechange',
                userid: userID,
            });
        });

        $('#rejectusernamerequest').on('click', function () {
            socket.Send({
                command: 'user.rejectusernamechange',
                userid: userID,
            });
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
            passwordTimeout = setTimeout(() => {
                socket.Send({
                    command: 'user.admin.edituser.changepassword',
                    userid: userID,
                    password: $(this).val(),
                });
            }, 1000 * timer);
        });

        $('[name="disabled"]').on('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changedisabled',
                userid: userID,
                checked: $(this).prop('checked'),
            });
        });

        $('[name="isadmin"]').on('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changeisadmin',
                userid: userID,
                checked: $(this).prop('checked'),
            });
        });

        $('[name="icon"]').on('change', function () {
            var reader = new FileReader();
            reader.onloadend = function () {
                socket.Send({
                    command: 'user.admin.edituser.uploadiconbase64',
                    userid: userID,
                    icon: reader.result,
                });
            }
            reader.readAsDataURL(this.files[0]);
        });

        $('[name="clearicon"]').on('click', function () {
            if (confirm('Are you Sure you want to clear the icon')) {
                socket.Send({
                    command: 'user.admin.edituser.clearicon',
                    userid: userID,
                });
            }
        });

        $('[name="delete"]').on('click', function () {
            if (confirm('Are you Sure you want to delete this user')) {
                socket.Send({
                    command: 'user.admin.edituser.deleteuser',
                    userid: userID,
                });
            }
        });



        $('[name="groups[]"]').on('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changegroup',
                userid: userID,
                checked: $(this).prop('checked'),
                group: $(this).data('group'),
            });
        });

        $('[name="permissions[]"]').on('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changepermission',
                userid: userID,
                checked: $(this).prop('checked'),
                permission: $(this).data('permission'),
            });
        });
    }

    function AcceptUsernameChange(data) {
        console.log(data);
        $('#username').val(data['name']);
        $('#usernamerequest').remove();
    }
    function RejectUsernameChange(data) {
        $('#usernamerequest').remove();
    }
    function UploadIconBase64(data) {
        $('#usericon').attr('src', data['name']);
    }
    function ClearIcon(data) {
        $('#usericon').attr('src', '');
        $('[name="icon"]').val(null);
        $('#updatedat').html(data['updatedat']);
    }

    function UpdateTimestamp(data) {
        console.log(data);
        if ('data' in data && 'updatedat' in data['data']) {
            let $uat = $('#updatedat').html(data['data']['updatedat']);
        }
    }
})();
