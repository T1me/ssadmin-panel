<html>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=0.64, user-scalable=no">
        <title>Panel</title>
        <style type="text/css">
            body
            {
                margin: 0px;
            }
            .sideBar
            {
                float: left;
                height: 100%;
                width: 80px;
                border-right: 10px;
                border-right-style: solid;
                border-right-color: #444444;
            }
            .sideBarText
            {
                position: fixed;
                left: 5px; 
                top: 20px
            }
            .main
            {
                float: left;
                margin-left: 20px;
                margin-top: 50px;
            }
            table
            {
                border-collapse: collapse;
            }
            th, td
            {
                border-style: solid;
                border-width: 1px;
                padding-left: 5px;
                padding-right: 5px;
            }
            th
            {
                background-color: #444444;
                border-color: black;
                color: white;
            }
        </style>
    </head>
    <body>
        <div class="sideBar">
            <div class="sideBarText">
                <p><a href="/panel">PANEL</a></p>
                <p><a href="/logout">LOGOUT</a></p>
            </div>
        </div>
        <div class="main">
            <table>
                <tr>
                    <th>Port</th>
                    <th>Password</th>
                    <th>Limit</th>
                    <th>Used</th>
                    <th>Remaining</th>
                </tr>
                {{range .}}
                <tr>
                    <td>{{.Port}}</td>
                    <td>{{.Password}}</td>
                    <td>{{.Limit}}</td>
                    <td>{{.Used}}</td>
                    <td>{{.Remaining}}</td>
                </tr>
                {{end}}
            </table>
        </div>
    </body>
</html>