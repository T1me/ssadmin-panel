<html>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
        <title>Login</title>
        <style type="text/css">
            .loginBox
            {
                border-style: double;
                max-width: 370px;
                min-height: 350px;
                margin: 130px auto 0;
                text-align: center;
                overflow-x: hidden;
            }
            @media screen and (max-width: 700px)
            {
                .loginBox
                {
                    margin: 0 auto;
                }
            }

            .prompt
            {
                height: 50px;
                text-align: center;
                font-size: 8px;
                color: red;
            }

            input
            {
                height: 40px;
                width: 300px;
            }

            .textInput
            {
                padding-left: 10px;
            }

            .button
            {
                font-weight: bold;
            }
        </style>
    </head>
    <body ondragstart="return false;" oncontextmenu="return false;" onselectstart="return false;">
        <div class="loginBox">
            <h1>SSADMIN-PANEL</h1>
            <div class="prompt">
                <br />
                <p>{{.}}</p>
                <br />
            </div>
            <form action="/" method="post">
                <div>
                    <p><input class="textInput" type="text" name="username" placeholder="Username" autocomplete="on"></p>
                </div>
                <div>
                    <p><input class="textInput" type="password" name="password" placeholder="Password" autocomplete="on"></p>
                </div>
                <div>
                    <p><input class="button" type="submit" value="Login"></p>
                </div>
            </form>
        </div>
        <!--
        <div style="position:absolute; right:0px; bottom:0px;">
            <img src="Gordon.png" alt="Gordon" />
        </div>
        -->
    </body>
</html>