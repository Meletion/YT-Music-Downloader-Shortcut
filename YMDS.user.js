// ==UserScript==
// @name         Music Downloader
// @namespace    https://meletion.github.io/bio
// @version      1.0.3
// @description  Download videos with ease
// @author       Meletion
// @match        https://www.youtube.com/watch?v=*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=youtube.com
// @grant        none
// ==/UserScript==
(function () {
    'use strict';
    function isTypingTarget(el) {
    const tag = el.nodeName;
    return (
        tag === "INPUT" ||
        tag === "TEXTAREA" ||
        tag === "SELECT" ||
        el.isContentEditable // catches <div contenteditable="true"> like YouTube's search
      );
    }
    document.addEventListener('keydown', function (e) {
        e = e || window.event;
        // Shift + D
        if (e.key == "D") {
            if (isTypingTarget(document.activeElement)) {
                return;
            }
            fetch('http://localhost:3000/download', {
                method: 'POST',
                headers: {
                    'Content-Type': 'text/plain'
                },
                body: new URLSearchParams(document.location.search).get('v')
            });
        }
    });
})();
