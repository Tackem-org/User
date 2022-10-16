(function () {
    let timer = 5; //seconds
    let usernameTimeout = null;
    let passwordTimeout = null;
    let userID = parseInt(window.location.href.split("/").pop());

    var ready = (callback) => {
        if (document.readyState != "loading") callback();
        else document.addEventListener("DOMContentLoaded", callback);
    };

    ready(() => {
        socket.AddReturnFunction(UpdateTimestamp);
        socket.AddReturnAction("user.acceptusernamechange", AcceptUsernameChange);
        socket.AddReturnAction("user.rejectusernamechange", RejectUsernameChange);
        socket.AddReturnAction("user.admin.edituser.uploadiconbase64", UploadIconBase64);
        socket.AddReturnAction("user.admin.edituser.clearicon", ClearIcon);


        document.getElementById('username').addEventListener('input', function () {
            if (this.value == "") {
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
                    username: this.value
                });
            }, 1000 * timer);
        });

        let acceptusernamerequest = document.getElementById('acceptusernamerequest');
        if (acceptusernamerequest != null) {
            acceptusernamerequest.addEventListener('click', function () {
                socket.Send({
                    command: 'user.acceptusernamechange',
                    userid: userID,
                });
            });
        }

        let rejectusernamerequest = document.getElementById('rejectusernamerequest')
        if (rejectusernamerequest != null) {

            rejectusernamerequest.addEventListener('click', function () {
                socket.Send({
                    command: 'user.rejectusernamechange',
                    userid: userID,
                });
            });
        }

        document.getElementById('password').addEventListener('input', function () {
            if (this.value == "") {
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
                    password: this.value,
                });
            }, 1000 * timer);
        });

        document.getElementById('disabled').addEventListener('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changedisabled',
                userid: userID,
                checked: this.checked,
            });
        });

        document.getElementById('isadmin').addEventListener('click', function () {
            socket.Send({
                command: 'user.admin.edituser.changeisadmin',
                userid: userID,
                checked: this.checked,
            });
        });

        document.getElementById('icon').addEventListener('change', function () {
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

        document.getElementById('clearicon').addEventListener('click', function () {
            if (confirm('Are you Sure you want to clear the icon')) {
                socket.Send({
                    command: 'user.admin.edituser.clearicon',
                    userid: userID,
                });
            }
        });

        document.getElementById('delete').addEventListener('click', function () {
            if (confirm('Are you Sure you want to delete this user')) {
                socket.Send({
                    command: 'user.admin.edituser.deleteuser',
                    userid: userID,
                });
            }
        });

        document.querySelectorAll('[name="groups[]"]').forEach((e) => {
            e.addEventListener("click", function () {
                socket.Send({
                    command: 'user.admin.edituser.changegroup',
                    userid: userID,
                    checked: this.checked,
                    group: parseInt(this.getAttribute('data-group')),
                });
            });
        });

        document.querySelectorAll('[name="permissions[]"]').forEach((e) => {
            e.addEventListener("click", function () {
                socket.Send({
                    command: 'user.admin.edituser.changepermission',
                    userid: userID,
                    checked: this.checked,
                    permission: parseInt(this.getAttribute('data-permission')),
                });
            });
        });
    });

    function AcceptUsernameChange(data) {
        document.getElementById('username').value = data['data']['name'];
        document.getElementById('usernamerequest').remove();
    }
    function RejectUsernameChange(data) {
        document.getElementById('usernamerequest').remove();
    }
    function UploadIconBase64(data) {
        document.getElementById('usericon').setAttribute('src', data['data']['icon']);
    }
    function ClearIcon(data) {
        document.getElementById('usericon').removeAttribute('src');
        document.querySelector('[name="icon"]').value = null;
    }

    function UpdateTimestamp(data) {
        if ('data' in data && 'updatedat' in data['data']) {
            document.getElementById('updatedat').innerHTML = data['data']['updatedat'];
        }
    }
})();
