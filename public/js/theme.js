var toggle = document.getElementById("theme-toggle");
var storedTheme = localStorage.getItem("theme") || (window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light");
var storedIcon = localStorage.getItem("theme-icon") || (storedTheme === "dark" ? "☀️" : "🌙");
if (storedTheme) {
    document.documentElement.setAttribute("data-theme", storedTheme);
    toggle.innerHTML = storedIcon;
}

toggle.onclick = function() {
    var currentTheme = document.documentElement.getAttribute("data-theme");
    var targetTheme = "light";
    var targetIcon = "🌙";

    if (currentTheme === "light") {
        targetTheme = "dark";
        targetIcon = "☀️";
    }

    document.documentElement.setAttribute("data-theme", targetTheme);
    toggle.innerHTML = targetIcon;
    localStorage.setItem("theme", targetTheme);
    localStorage.setItem("theme-icon", targetIcon);
}