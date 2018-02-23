<html>
    <head>
        <title>Login</title>
        <style type="text/css">
            .loginBox
            {
                width: 280px;
                height: 100px;
                position: fixed;
                left: 50%;
                top: 30%;
                z-index: 11;
                margin: -50px 0 0 -140px;
                text-align: center;
                font-size: 0.9em;
            }
            .prompt
            {
                width: 280px;
                height: 30px;
                position: fixed;
                left: 50%;
                top: 20%;
                z-index: 11;
                margin: -50px 0 0 -140px;
                text-align: center;
                font-size: 0.9em;
                color: red;
            }
        </style>
    </head>
    <body ondragstart="return false;" oncontextmenu="return false;" onselectstart="return false;">
        <div class="prompt">
            <p>{{.}}</p>
        </div>
        <div class="loginBox">
            <form action="/" method="post">
                <div>
                    Username：<input type="text" name="username">
                </div>
                <div>
                    Password ：<input type="password" name="password">
                </div>
                <p><input type="submit" value="Login"></p>
            </form>
        </div>
        <!--
        <div style="position:absolute; right:0px; bottom:0px;">
            <img src="Gordon.png" alt="Gordon" />
        </div>
        -->
    </body>
</html>