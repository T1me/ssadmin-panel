<html>
    <head>
        <title>Panel</title>
        <style type="text/css">
            .sideBar
            {
                position: absolute;
                left: 0;
                top: 0;
                bottom: 0;
                width: 100px;
                border-right: 10px;
                border-right-style: solid;
                border-right-color: #444444;
            }
            .sideBarText
            {
                position: fixed;
                left: 10px; 
            }
            .table
            {
                position: fixed;
                left: 30%;
            }
        </style>
    </head>
        <div class="sideBar">
            <div class="sideBarText">
                <p><a href="/panel">PANEL</a></p>
                <p><a href="/logout">LOGOUT</a></p>
            </div>
        </div>
        <div class="table">
            <table border="1">
                <tr>
                    <th>Port</th>
                    <th>Password</th>
                    <th>Limit</th>
                    <th>Used</th>
                    <th>Remaining</th>
                </tr>
                {{range .}}
                <tr>
                    <th>{{.Port}}</th>
                    <th>{{.Password}}</th>
                    <th>{{.Limit}}</th>
                    <th>{{.Used}}</th>
                    <th>{{.Remaining}}</th>
                </tr>
                {{end}}
            </table>
        </div>
</html>