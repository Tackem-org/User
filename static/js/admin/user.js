(function () {
    let socket;

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

        ChangeUsername() {
            this.#ws.send(JSON.stringify({
                command: 'changeusername',
                userid: this.parseUserID(),

            }));
        }

        UpdatePassword() {
            this.#ws.send(JSON.stringify({
                command: 'updatepassword',
                userid: this.parseUserID(),

            }));
        }

        ChangeDisabled() {
            this.#ws.send(JSON.stringify({
                command: 'changedisabled',
                userid: this.parseUserID(),

            }));
        }

        ChangeIsAdmin() {
            this.#ws.send(JSON.stringify({
                command: 'changeisadmin',
                userid: this.parseUserID(),

            }));
        }
        DeleteUser() {
            this.#ws.send(JSON.stringify({
                command: 'deleteuser',
                userid: this.parseUserID(),

            }));
        }
        ChangeGroup() {
            this.#ws.send(JSON.stringify({
                command: 'changegroup',
                userid: this.parseUserID(),

            }));
        }

        ChangePermission() {
            this.#ws.send(JSON.stringify({
                command: 'changepermission',
                userid: this.parseUserID(),

            }));
        }


        MessageListener(e) {
            let edata = JSON.parse(e.data);
            let data = edata['data'];
            switch (data['command']) {
                case 'changeusername':
                case 'updatepassword':
                case 'changedisabled':
                case 'changeisadmin':
                case 'deleteuser':
                case 'setgroup':
                case "setpermission":
                default:
                    console.log('ERROR COMMAND NOT KNOWN SKIPPING:', e.data);
                    break;
            }

        };
    }
})();
