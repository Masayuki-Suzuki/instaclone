<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <!-- <link rel="icon" href="/favicon.ico"/> -->
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <meta
            name="description"
            content="{{.Description}}"
    />
    <!--
    <link rel="apple-touch-icon" href="%PUBLIC_URL%/logo192.png"/>
    -->
    <!--
      manifest.json provides metadata used when your web app is installed on a
      user's mobile device or desktop. See https://developers.google.com/web/fundamentals/web-app-manifest/
      <link rel="manifest" href="%PUBLIC_URL%/manifest.json" />
    -->
    <title>{{.Title}}</title>
</head>
<body style="font-family:sans-serif; padding: 24px;">
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <h1>{{.PageTitle}}</h1>
    <div>
        {{range .Users}}
            <div style="border-bottom: solid 1px #ccc;">
                <h2 style="font-size: 20px">{{.FullName}}</h2>
                <p>ID: {{.Uid}}</p>
                <p>User Name: {{.Username}}</p>
                <p>Email: {{.Email}}</p>
            </div>
        {{end}}
    </div>
    <div id="root"></div>
    <script src="{{.ReactFilePath}}"></script>
</body>
</html>
