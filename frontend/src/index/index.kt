package index

import app.*
import kotlinext.js.*
import react.dom.*
import kotlin.browser.*

fun main(args: Array<String>) {
    requireAll(require.context("src", true, js("/\\.css$/")))
    requireAll(require.context("src", true, js("/\\.scss$/")))
    require("react-select/dist/react-select.css")
    window.onload = {
        render(document.getElementById("root")) {
            app()
        }
    }

}
