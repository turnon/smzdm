package templates

func init() {
	Templates["templates/rich"] = `<!DOCTYPE html>
<html>

<head>
    <meta charset="UTF-8" />
    <meta name="referrer" content="no-referrer" />
    <title>{{ .now }}</title>
    <style>
        a {
            text-decoration: none;
            color: black;
        }

        nav {
            width: 8em;
            font-size: 1.1em;
            position: fixed;
            top: 0;
            left: 0;
            bottom: 0;
            overflow-y: auto;
            padding: 0;
            background-color: #ff1966;
        }

        .searches {
            padding: 0 3em 2em 11em;
        }

        nav a {
            display: block;
            color: white;
            padding: .3em 0 0 1.2em;
        }

        nav a:visited {
            color: white;
            text-decoration: none;
        }

        nav a:hover {
            color: #FFEB3B;
        }

        nav a:first-child {
            padding-top: 1em;
        }

        .entry {
            display: inline-block;
            width: 120px;
            height: 200px;
        }

        img {
            max-width: 100px;
            max-height: 100px;
        }

        .info {
            display: inline;
        }

        .info p {
            width: 100px;
            display: inline-block;
            margin: 0;
        }

        h1 {
            border-bottom: 1px dotted #ff1966;
            margin-top: 0;
            padding: 0 0 10px 0;
            color: #ff1966;
            font-weight: normal;
        }

        .time,
        .title {
            color: #616161;
        }

        .price {
            color: #e91e63;
        }
    </style>
</head>

<body>
    <nav>
        {{ range .data }}
        <a class='nav' href='#{{ .Keyword }}'>
            {{ .Keyword }}
        </a>
        {{ end }}
    </nav>
    {{ range .data }}
    <div id='{{ .Keyword }}' class='searches'>
        <h1 class='keyword'> {{ .Keyword }} </h1>
        <div class='entries'>
            {{ range .Entries }}
            <a href='{{ .Href }}' target='_blank'>
                <div class='entry'>
                    <img src='{{ .Img }}' />
                    <div class='info'>
                        <p class='time'> {{ .Time }} </p>
                        <p class='title'> {{ .Title }} </p>
                        <p class='price'> {{ .Price }} </p>
                    </div>
                </div>
            </a>
            {{ end }}
        </div>

    </div>
    {{ end }}
</body>

</html>`
}
