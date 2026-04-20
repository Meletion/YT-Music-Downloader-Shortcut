// ==UserScript==
// @name         Music Downloader
// @namespace    https://meletion.github.io/bio
// @version      1
// @description  Download videos with ease
// @author       Meletion
// @match        https://www.youtube.com/watch?v=*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=youtube.com
// @grant        none
// ==/UserScript==

(function () {
    'use strict';

    document.addEventListener('keydown', function (e) {
        e = e || window.event;

        // Shift + d
        if (e.key == "D") {
            fetch('http://localhost:3000/download', {
                method: 'POST',
                headers: {
                    'Content-Type': 'text/plain'
                },
                body: new URLSearchParams(document.location.search).get('v')
            });

        }

    }) 

})();
