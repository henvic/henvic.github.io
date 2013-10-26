/*jslint browser: true, nomen: false */

(function () {
    'use strict';

    var headerRowH1 = document.querySelector('.header-row h1'),
        blockquoteOnBottom = document.querySelector('.blockquote-on-bottom'),
        blockquoteOnBottomListener,
        scrollHeaderListener,
        hasVanilla = false;

    scrollHeaderListener = function () {
        if (window.pageYOffset > 50 && !hasVanilla) {
            headerRowH1.classList.add('vanilla');
            hasVanilla = true;
        }

        if (window.pageYOffset <= 50 && hasVanilla) {
            headerRowH1.classList.remove('vanilla');
            hasVanilla = false;
        }
    };

    window.addEventListener('scroll', scrollHeaderListener);

    blockquoteOnBottomListener = function () {
        blockquoteOnBottom.classList.add('blockquote-on-bottom-alive');
        setTimeout(function () {
            blockquoteOnBottom.classList.remove('blockquote-on-bottom-alive');
        }, 2000);
        blockquoteOnBottom.removeEventListener('mouseover', blockquoteOnBottomListener);
    };

    blockquoteOnBottom.addEventListener('mouseover', blockquoteOnBottomListener);

    /**
     * Twitter follow button code
     */
    (function (d, s, id) {
        var js,
            fjs = d.getElementsByTagName(s)[0],
            p = /^http:/.test(d.location) ? 'http' : 'https';

        if (!d.getElementById(id)) {
            js = d.createElement(s);
            js.id = id;
            js.src = p + '://platform.twitter.com/widgets.js';
            fjs.parentNode.insertBefore(js, fjs);
        }
    }(document, 'script', 'twitter-wjs'));
}());
