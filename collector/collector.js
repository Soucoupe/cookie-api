t = {
    PLUGINS: ['WebEx64 General Plugin Container', 'YouTube Plug-in', 'Java Applet Plug-in', 'Shockwave Flash', 'iPhotoPhotocast', 'SharePoint Browser Plug-in', 'Chrome Remote Desktop Viewer', 'Chrome PDF Viewer', 'Native Client', 'Unity Player', 'WebKit-integrierte PDF', 'QuickTime Plug-in', 'RealPlayer Version Plugin', 'RealPlayer(tm) G2 LiveConnect-Enabled Plug-In (32-bit)', 'Mozilla Default Plug-in', 'Adobe Acrobat', 'AdobeAAMDetect', 'Google Earth Plug-in', 'Java Plug-in 2 for NPAPI Browsers', 'Widevine Content Decryption Module', 'Microsoft Office Live Plug-in', 'Windows Media Player Plug-in Dynamic Link Library', 'Google Talk Plugin Video Renderer', 'Edge PDF Viewer', 'Shockwave for Director', 'Default Browser Helper', 'Silverlight Plug-In'],
    ins: null,
}

function run() {
    if (window.speechSynthesis) {

        var a = window.speechSynthesis.getVoices();

        a = window.speechSynthesis.getVoices();

        a = window.speechSynthesis.getVoices();
    }

    var window_model = {

        'screenData': {

            'availWidth': window['screen']['availWidth'],

            'availHeight': window['screen']['availHeight'],

            'width': window['screen']['width'],

            'height': window['screen']['height'],

            'colorDepth': window['screen']['colorDepth'],

            'pixelDepth': window['screen']['pixelDepth'],



            'innerWidth': window['innerWidth'] || document['body']['clientWidth'],

            'innerHeight': window['innerHeight'] || document['body']['clientHeight'],

            'outerWidth': window['outerWidth'] || document['body']['outerWidth'],

        },

        'navigator': {

            'userAgent': window.navigator.userAgent,

            'productSub': navigator.productSub,

            'language': navigator.language,

            'product': navigator.product,

            'onLine': navigator.onLine,

            'vibrate': 'function' == typeof navigator['vibrate'],

            'getBattery': 'function' == typeof navigator['getBattery'],

            'credentials': Boolean(navigator['credentials']),

            'appMinorVersion': Boolean(navigator['appMinorVersion']),

            'bluetooth': Boolean(navigator['bluetooth']),

            'storage': Boolean(navigator['storage']),

            'getGamepads': Boolean(navigator['getGamepads']),

            'getStorageUpdates': Boolean(navigator['getStorageUpdates']),

            'hardwareConcurrency': Boolean(navigator['hardwareConcurrency']),

            'mediaDevices': Boolean(navigator['mediaDevices']),

            'mozAlarms': Boolean(navigator['mozAlarms']),

            'mozConnection': Boolean(navigator['mozConnection']),

            'mozIsLocallyAvailable': Boolean(navigator['mozIsLocallyAvailable']),

            'mozPhoneNumberService': Boolean(navigator['mozPhoneNumberService']),

            'msManipulationViewsEnabled': Boolean(navigator['msManipulationViewsEnabled']),

            'permissions': Boolean(navigator['permissions']),

            'registerProtocolHandler': Boolean(navigator['registerProtocolHandler']),

            'requestMediaKeySystemAccess': Boolean(navigator['requestMediaKeySystemAccess']),

            'requestWakeLock': Boolean(navigator['requestWakeLock']),

            'sendBeacon': Boolean(navigator['sendBeacon']),

            'serviceWorker': Boolean(navigator['serviceWorker']),

            'storeWebWideTrackingException': Boolean(navigator['storeWebWideTrackingException']),

            'webkitGetGamepads': Boolean(navigator['webkitGetGamepads']),

            'webkitTemporaryStorage': Boolean(navigator['webkitTemporaryStorage']),

            'webdriver': navigator['webdriver'],

            'cookieEnabled': navigator['cookieEnabled'] ? navigator['cookieEnabled'] : -1,

            'javaEnabled': navigator['javaEnabled'] ? navigator['javaEnabled']() : -1,

            'doNotTrack': navigator['doNotTrack'] ? navigator['doNotTrack'] : -1,

        },

        'addEventListener': Boolean(window['addEventListener']),

        'XMLHttpRequest': Boolean(window['XMLHttpRequest']),

        'XDomainRequest': Boolean(window['XDomainRequest']),

        'emit': window['emit'],

        'DeviceOrientationEvent': Boolean(window['DeviceOrientationEvent']),

        'DeviceMotionEvent': Boolean(window['DeviceMotionEvent']),

        'TouchEvent': Boolean(window['TouchEvent']),

        'spawn': window['spawn'],

        'chrome': window['chrome'],

        'prototype_bind': Boolean(Function['prototype']['bind']),

        'Buffer': window['Buffer'],

        'PointerEvent': Boolean(window['PointerEvent']),

        '_phantom': window['_phantom'],

        'webdriver': window['webdriver'],

        'domAutomation': window['domAutomation'],

        'callPhantom': window['callPhantom'],

        'CC_ON': new Function('return/*@cc_on!@*/!1')(),

        'document': {

            'documentMode': typeof document['documentMode'],

            'webdriver': window['document']['documentElement']['getAttribute']('webdriver') != null,

            'driver': window['document']['documentElement']['getAttribute']('driver') != null,

            'selenium': null != window['document']['documentElement']['getAttribute']('selenium'),

            'hidden': document['hidden'],

            'mozHidden': document['mozHidden'],

            'msHidden': document['msHidden'],

            'webkitHidden': document['webkitHidden']

        },

        'opera': window['opera'],

        'InstallTrigger': 'undefined' != typeof InstallTrigger,

        'HTMLElement': window['HTMLElement'] && Object['prototype']['toString']['call'](window['HTMLElement'])[

            'indexOf']('Constructor') > 0,

        'webRTC': 'function' == typeof window['RTCPeerConnection'] || 'function' == typeof window[

            'mozRTCPeerConnection'] || 'function' == typeof window['webkitRTCPeerConnection'],

        'mozInnerScreenY': 'mozInnerScreenY' in window ? window['mozInnerScreenY'] : 0,

        'prototype_forEach': Boolean(Array['prototype']['forEach']),

        'FileReader': 'FileReader' in window,

        'imul': Boolean(Math['imul']),

        'parseInt': Boolean(Number['parseInt']),

        'hypot': Boolean(Math['hypot']),

        'value1': Boolean(window['$cdc_asdjflasutopfhvcZLmcfl_'] || document[

            '$cdc_asdjflasutopfhvcZLmcfl_']),

        'XPathResult': void 0 !== window['XPathResult'] || void 0 !== document['XPathResult'],

        'sessionStorage': !!window['sessionStorage'],

        'localStorage': !!window['localStorage'],

        'indexedDB': !!window['indexedDB'],

        'px400': _PX400()

        //'csh': window['ssh']

    };



    var plugins = [];

    for (var i = 0; i < navigator['plugins'].length; i++) {

        plugins.push(navigator['plugins'][i].name);

    }



    window_model['navigator']['plugins'] = plugins;



    var rCFP = [];



    function canvas(t) {

        var a = -1;

        var e = document['createElement']('canvas');

        e['width'] = 280;

        e['height'] = 60;

        e['style']['display'] = 'none';

        if ('function' == typeof e['getContext']) {

            var n = e['getContext']('2d');

            n['fillStyle'] = 'rgb(102, 204, 0)';

            n['fillRect'](100, 5, 80, 50);

            n['fillStyle'] = '#f60';

            n['font'] = '16pt Arial';

            n['fillText'](t, 10, 40);

            n['strokeStyle'] = 'rgb(120, 186, 176)';

            n['arc'](80, 10, 20, 0, Math['PI'], !1);

            n['stroke']();



            var o = e['toDataURL']();



            a = 0;

            for (var m = 0; m < o['length']; m++) {

                a = (a << 5) - a + o['charCodeAt'](m), a &= a

            }



            a = a['toString']();



            for (var rVal = 0; rVal <= 1000; rVal++) {

                var r = document['createElement']('canvas');



                r['width'] = 16;

                r['height'] = 16;



                var i = r['getContext']('2d');



                i['font'] = '6pt Arial';

                i['fillText'](rVal, 1, 12);

                var c = r['toDataURL'](),

                    b = 0;

                for (var d = 0; d < c['length']; d++) {

                    b = (b << 5) - b + c['charCodeAt'](d), b &= b

                }



                rCFP[rVal] = b['toString']()

            }



            return a;

        }

    }



    // window_model['canvas'] = canvas("<@nv45. F1n63r,Pr1n71n6!");



    window_model['canvas'] = {

        'value1': canvas("<@nv45. F1n63r,Pr1n71n6!"),

        'value2': canvas("m,Ev!xV67BaU> eh2m<f3AG3@")

    }

    window_model['rCFP'] = rCFP;



    function get_cf_date() {

        return Date['now'] ? Date['now']() : +new Date

    }



    function fonts_optm() {

        var a = 200,

            e = get_cf_date(),

            n = [];

        var o = ['sans-serif', 'monospace'],

            m = [0, 0],

            r = [0, 0],

            i = document['createElement']('div');

        i['style']['cssText'] = 'position: relative; left: -9999px; visibility: hidden; display: block !important';

        var c;

        for (c = 0; c < o['length']; c++) {

            var b = document['createElement']('span');

            b['innerHTML'] = 'abcdefhijklmnopqrstuvxyz1234567890;+-.';

            b['style']['fontSize'] = '90px';

            b['style']['fontFamily'] = o[c];

            i['appendChild'](b);

        }

        for (document['body']['appendChild'](i), c = 0; c <

            i['childNodes']['length']; c++) {

            b = i['childNodes'][c];

            m[c] = b['offsetWidth'];

            r[c] = b['offsetHeight'];

        }

        document['body']['removeChild'](i);



        if (get_cf_date() - e > a)

            return n;



        var d = ['Geneva', 'Lobster', 'New York', 'Century', 'Apple Gothic', 'Minion Pro', 'Apple LiGothic',

            'Century Gothic', 'Monaco', 'Lato', 'Fantasque Sans Mono', 'Adobe Braille', 'Cambria', 'Futura',

            'Bell MT', 'Courier', 'Courier New', 'Calibri', 'Avenir Next', 'Birch Std', 'Palatino',

            'Ubuntu Regular', 'Oswald', 'Batang', 'Ubuntu Medium', 'Cantarell', 'Droid Serif', 'Roboto',

            'Helvetica Neue', 'Corsiva Hebrew', 'Adobe Hebrew', 'TI-Nspire', 'Comic Neue', 'Noto', 'AlNile',

            'Palatino-Bold', 'ArialHebrew-Light', 'Avenir', 'Papyrus', 'Open Sans', 'Times', 'Quicksand',

            'Source Sans Pro', 'Damascus', 'Microsoft Sans Serif'

        ],

            k = document['createElement']('div');

        k['style']['cssText'] = 'position: relative; left: -9999px; visibility: hidden; display: block !important';



        for (var l = [], s = 0; s < d['length']; s++) {

            var u = document['createElement']('div');



            for (c = 0; c < o['length']; c++) {

                var b = document['createElement']('span');

                b['innerHTML'] = 'abcdefhijklmnopqrstuvxyz1234567890;+-.';

                b['style']['fontSize'] = '90px';

                b['style']['fontFamily'] = d[s] + ',' + o[c];

                u['appendChild'](b)

            }



            k['appendChild'](u)

        }



        if (get_cf_date() - e > a) {

            return n;

        }



        document['body']['appendChild'](k);



        for (var s = 0; s < k['childNodes']['length']; s++) {

            var _ = !1,

                u = k['childNodes'][s];



            for (c = 0; c < u['childNodes']['length']; c++) {

                var b = u['childNodes'][c];

                if (b['offsetWidth'] !== m[c] || b['offsetHeight'] !== r[c]) {

                    _ = !0;

                    break

                }

            }



            if (_ && l['push'](s), get_cf_date() - e > a) break

        }



        document['body']['removeChild'](k);

        n = l['sort']()

        return n;

    }



    function fonts() {

        var a = [];

        var e = ['serif', 'sans-serif', 'monospace'],

            n = [0, 0, 0],

            o = [0, 0, 0],

            m = document['createElement']('span');



        m['innerHTML'] = 'abcdefhijklmnopqrstuvxyz1234567890;+-.';

        m['style']['fontSize'] = '90px';



        var r;

        for (r = 0; r < e['length']; r++) {

            m['style']['fontFamily'] = e[r];

            document['body']['appendChild'](m);

            n[r] = m['offsetWidth'];

            o[r] = m['offsetHeight'];

            document['body']['removeChild'](m);

        }



        for (var i = ['Geneva', 'Lobster', 'New York', 'Century', 'Apple Gothic', 'Minion Pro', 'Apple LiGothic',

            'Century Gothic', 'Monaco', 'Lato', 'Fantasque Sans Mono', 'Adobe Braille', 'Cambria', 'Futura',

            'Bell MT', 'Courier', 'Courier New', 'Calibri', 'Avenir Next', 'Birch Std', 'Palatino',

            'Ubuntu Regular', 'Oswald', 'Batang', 'Ubuntu Medium', 'Cantarell', 'Droid Serif', 'Roboto',

            'Helvetica Neue', 'Corsiva Hebrew', 'Adobe Hebrew', 'TI-Nspire', 'Comic Neue', 'Noto', 'AlNile',

            'Palatino-Bold', 'ArialHebrew-Light', 'Avenir', 'Papyrus', 'Open Sans', 'Times', 'Quicksand',

            'Source Sans Pro', 'Damascus', 'Microsoft Sans Serif'

        ], c = [], b = 0; b < i['length']; b++) {

            var d = !1;



            for (r = 0; r < e['length']; r++)

                if (m['style']['fontFamily'] = i[b] + ',' + e[r], document['body']['appendChild'](m), m[

                    'offsetWidth'] === n[r] && m['offsetHeight'] === o[r] || (d = !0), document['body'][

                        'removeChild'

                    ](

                        m), d) {

                    c['push'](b);

                    break

                }

        }

        a = c['sort']()

        return a;

    }



    window_model['fonts_optm'] = fonts_optm();

    window_model['fonts'] = fonts();



    function getmr() {

        try {

            if ('undefined' == typeof performance || void 0 === performance['now'] || 'undefined' == typeof JSON)

                return void (window_model['mr'] = 'undef');

            for (var a = '', t = 1e3, e = [Math['abs'], Math['acos'], Math['asin'], Math['atanh'], Math['cbrt'], Math[

                'exp'], Math['random'], Math['round'], Math['sqrt'], isFinite, isNaN, parseFloat, parseInt,

            JSON['parse']

            ], n = 0; n < e['length']; n++) {

                var o = [],

                    m = 0,

                    r = performance['now'](),

                    i = 0,

                    c = 0;

                if (void 0 !== e[n]) {

                    for (i = 0; i < t && m < .6; i++) {

                        for (var b = performance['now'](), d = 0; d < 4e3; d++) e[n](3.14);

                        var k = performance['now']();

                        o['push'](Math['round'](1e3 * (k - b))), m = k - r

                    }

                    var l = o['sort']();

                    c = l[Math['floor'](l['length'] / 2)] / 5

                }

                a = a + c + ','

            }

            window_model['mr'] = a

        } catch (a) {

            window_model['mr'] = 'exception'

        }

    }

    function csh() {

        /*

csh: function () {

if (window.speechSynthesis) {

    var a = window.speechSynthesis.getVoices();

    if (a.length > 0) {

        for (var t = '', e = 0; e < a.length; e++) t += a[e].voiceURI + '_' + a[e].lang;

        ssh = ats(mn_s(t))

    } else ssh = '0'

} else ssh = 'n'

},*/

        if (window.speechSynthesis) {

            var a = window.speechSynthesis.getVoices();

            if (a.length > 0) {

                for (var t = '', e = 0; e < a.length; e++) t += a[e].voiceURI + '_' + a[e].lang;

                window_model['ssh'] = ats(mn_s(t))

            } else window_model['ssh'] = '0'

        } else window_model['ssh'] = 'n'



    }





    function rotate_right(a, t) {

        return a >>> t | a << 32 - t

    }

    function encode_utf8(a) {

        return unescape(encodeURIComponent(a))

    }

    function mn_s(a) {

        var t = [1116352408, 1899447441, 3049323471, 3921009573, 961987163, 1508970993, 2453635748, 2870763221, 3624381080, 310598401, 607225278, 1426881987, 1925078388, 2162078206, 2614888103, 3248222580, 3835390401, 4022224774, 264347078, 604807628, 770255983, 1249150122, 1555081692, 1996064986, 2554220882, 2821834349, 2952996808, 3210313671, 3336571891, 3584528711, 113926993, 338241895, 666307205, 773529912, 1294757372, 1396182291, 1695183700, 1986661051, 2177026350, 2456956037, 2730485921, 2820302411, 3259730800, 3345764771, 3516065817, 3600352804, 4094571909, 275423344, 430227734, 506948616, 659060556, 883997877, 958139571, 1322822218, 1537002063, 1747873779, 1955562222, 2024104815, 2227730452, 2361852424, 2428436474, 2756734187, 3204031479, 3329325298],

            e = 1779033703,

            n = 3144134277,

            o = 1013904242,

            m = 2773480762,

            r = 1359893119,

            i = 2600822924,

            c = 528734635,

            b = 1541459225,

            d = encode_utf8(a),

            k = 8 * d.length;

        d += String.fromCharCode(128);

        for (var s = d.length / 4 + 2, l = Math.ceil(s / 16), u = new Array(l), _ = 0; _ < l; _++) {

            u[_] = new Array(16);

            for (var f = 0; f < 16; f++) u[_][f] = d.charCodeAt(64 * _ + 4 * f) << 24 | d.charCodeAt(64 * _ + 4 * f + 1) << 16 | d.charCodeAt(64 * _ + 4 * f + 2) << 8 | d.charCodeAt(64 * _ + 4 * f + 3) << 0

        }

        var p = k / Math.pow(2, 32);

        u[l - 1][14] = Math.floor(p), u[l - 1][15] = k;

        for (var v = 0; v < l; v++) {

            for (var h, g = new Array(64), w = e, y = n, E = o, S = m, C = r, h = i, M = c, x = b, _ = 0; _ < 64; _++) {

                var j, A, L, P, T, D;

                _ < 16 ? g[_] = u[v][_] : (j = rotate_right(g[_ - 15], 7) ^ rotate_right(g[_ - 15], 18) ^ g[_ - 15] >>> 3, A = rotate_right(g[_ - 2], 17) ^ rotate_right(g[_ - 2], 19) ^ g[_ - 2] >>> 10, g[_] = g[_ - 16] + j + g[_ - 7] + A), A = rotate_right(C, 6) ^ rotate_right(C, 11) ^ rotate_right(C, 25), L = C & h ^ ~C & M, P = x + A + L + t[_] + g[_], j = rotate_right(w, 2) ^ rotate_right(w, 13) ^ rotate_right(w, 22), T = w & y ^ w & E ^ y & E, D = j + T, x = M, M = h, h = C, C = S + P >>> 0, S = E, E = y, y = w, w = P + D >>> 0

            }

            e += w, n += y, o += E, m += S, r += C, i += h, c += M, b += x

        }

        return [e >> 24 & 255, e >> 16 & 255, e >> 8 & 255, 255 & e, n >> 24 & 255, n >> 16 & 255, n >> 8 & 255, 255 & n, o >> 24 & 255, o >> 16 & 255, o >> 8 & 255, 255 & o, m >> 24 & 255, m >> 16 & 255, m >> 8 & 255, 255 & m, r >> 24 & 255, r >> 16 & 255, r >> 8 & 255, 255 & r, i >> 24 & 255, i >> 16 & 255, i >> 8 & 255, 255 & i, c >> 24 & 255, c >> 16 & 255, c >> 8 & 255, 255 & c, b >> 24 & 255, b >> 16 & 255, b >> 8 & 255, 255 & b]

    }









    /*

wgl: function () {

try {

    var a = document.createElement('canvas'),

        t = a.getContext('webgl');

    wv = 'n', wr = 'n', weh = 'n', wl = 0, t && (wv = 'b', wr = 'b', weh = 'b', t.getSupportedExtensions() && (weh = ats(mn_s(JSON.stringify(t.getSupportedExtensions().sort()))), wl = t.getSupportedExtensions().length, t.getSupportedExtensions().indexOf('WEBGL_debug_renderer_info') >= 0 && (wv = t.getParameter(t.getExtension('WEBGL_debug_renderer_info').UNMASKED_VENDOR_WEBGL), wr = t.getParameter(t.getExtension('WEBGL_debug_renderer_info').UNMASKED_RENDERER_WEBGL))))

} catch (a) {

    wv = 'e', wr = 'e', weh = 'e', wl = 0

}

},*/

    function wgl() {



        try {

            var a = document.createElement('canvas'),

                t = a.getContext('webgl');

            window_model['wv'] = 'n', window_model['wr'] = 'n', window_model['weh'] = 'n', window_model['wl'] = 0, t && (window_model['wv'] = 'b', window_model['wr'] = 'b', window_model['weh'] = 'b', t.getSupportedExtensions() && (window_model['weh'] = ats(mn_s(JSON.stringify(t.getSupportedExtensions().sort()))), window_model['wl'] = t.getSupportedExtensions().length, t.getSupportedExtensions().indexOf('WEBGL_debug_renderer_info') >= 0 && (window_model['wv'] = t.getParameter(t.getExtension('WEBGL_debug_renderer_info').UNMASKED_VENDOR_WEBGL), window_model['wr'] = t.getParameter(t.getExtension('WEBGL_debug_renderer_info').UNMASKED_RENDERER_WEBGL))))

        } catch (a) {

            window_model['wv'] = 'e', window_model['wr'] = 'e', window_model['weh'] = 'e', window_model['wl'] = 0

        }

    }







    function np() {

        var a = [],

            t = ['geolocation', 'notifications', 'push', 'midi', 'camera', 'microphone', 'speaker', 'device-info', 'background-sync', 'bluetooth', 'persistent-storage', 'ambient-light-sensor', 'accelerometer', 'gyroscope', 'magnetometer', 'clipboard', 'accessibility-events', 'clipboard-read', 'clipboard-write', 'payment-handler'];

        try {

            if (!navigator['permissions']) return 6;

            var e = function (t, e) {

                return navigator['permissions']['query']({

                    name: t

                })['then'](function (t) {

                    switch (t['state']) {

                        case 'prompt':

                            a[e] = 1;

                            break;

                        case 'granted':

                            a[e] = 2;

                            break;

                        case 'denied':

                            a[e] = 0;

                            break;

                        default:

                            a[e] = 5

                    }

                })['catch'](function (t) {

                    a[e] = -1 !== t['message']['indexOf']('is not a valid enum value of type PermissionName') ? 4 : 3

                })

            },

                n = t['map'](function (a, t) {

                    return e(a, t)

                });

            Promise['all'](n)['then'](function () {

                var N = a['join']('');

                window_model['np'] = N;

                load();

            })

        } catch (a) {

            return 7

        }

    }





    function ats(a) {

        for (var t = '', e = 0; e < a.length; e++) t += 2 == a[e].toString(16).length ? a[e].toString(16) : '0' + a[e].toString(16);

        return t

    }









    /*

m: function () {

var a = ['Monospace', 'Wingdings 2', 'ITC Bodoni 72 Bold', 'Menlo', 'Gill Sans MT', 'Lucida Sans', 'Bodoni 72', 'Serif', 'Shree Devanagari 714', 'Microsoft Tai Le', 'Nimbus Roman No 9 L', 'Candara', 'Press Start 2P', 'Waseem'],

    t = document.createElement('span');

t.innerHTML = 'mmmmmmmmlli', t.style.fontSize = '192px';

var e = '',

    n = document.getElementsByTagName('body')[0];

if (n) {

    for (var o in a) t.style.fontFamily = a[o], n.appendChild(t), e += a[o] + ':' + t.offsetWidth + ',' + t.offsetHeight + ';', n.removeChild(t);

    fmh = ats(mn_s(e))

} else fmh = '';

},*/

    fmz = 'devicePixelRatio' in window && void 0 !== window.devicePixelRatio ? window.devicePixelRatio : -1



    function fm() {

        var a = ['Monospace', 'Wingdings 2', 'ITC Bodoni 72 Bold', 'Menlo', 'Gill Sans MT', 'Lucida Sans', 'Bodoni 72', 'Serif', 'Shree Devanagari 714', 'Microsoft Tai Le', 'Nimbus Roman No 9 L', 'Candara', 'Press Start 2P', 'Waseem'],

            t = document.createElement('span');

        t.innerHTML = 'mmmmmmmmlli', t.style.fontSize = '192px';

        var e = '',

            n = document.getElementsByTagName('body')[0];

        if (n) {

            for (var o in a) t.style.fontFamily = a[o], n.appendChild(t), e += a[o] + ':' + t.offsetWidth + ',' + t.offsetHeight + ';', n.removeChild(t);

            window_model['fmh'] = ats(mn_s(e))

        } else fmh = '';

        window_model['fmz'] = 'devicePixelRatio' in window && void 0 !== window.devicePixelRatio ? window.devicePixelRatio : -1





    }


    function canvas(a) {
        try {
            var e = -1;
            if (!sf4()) {
                var n = document.createElement('canvas');
                if (n.width = 280, n.height = 60, n.style.display = 'none', 'function' == typeof n.getContext) {
                    var o = n.getContext('2d');
                    o.fillStyle = 'rgb(102, 204, 0)', o.fillRect(100, 5, 80, 50), o.fillStyle = '#f60', o.font = '16pt Arial', o.fillText(a, 10, 40), o.strokeStyle = 'rgb(120, 186, 176)', o.arc(80, 10, 20, 0, Math.PI, !1), o.stroke();
                    var m = n.toDataURL();
                    e = 0;
                    for (var r = 0; r < m.length; r++) {
                        e = (e << 5) - e + m.charCodeAt(r), e &= e
                    }
                    e = e.toString();
                    var i = document.createElement('canvas');
                    i.width = 16, i.height = 16;
                    var c = i.getContext('2d');
                    c.font = '6pt Arial', rVal = Math.floor(1e3 * Math.random()).toString(), c.fillText(rVal, 1, 12);
                    for (var b = i.toDataURL(), d = 0, k = 0; k < b.length; k++) {
                        d = (d << 5) - d + b.charCodeAt(k), d &= d
                    }
                    rCFP = d.toString()
                }
            }
            return e
        } catch (a) {
            console.log(a)
            return 'exception'
        }
    }

    function uar() {
        return window.navigator.userAgent.replace(/\\|"/g, '')
    }

    function sf4() {
        var a = uar();
        return !(!~a.indexOf('Version/4.0') || !(~a.indexOf('iPad;') || ~a.indexOf('iPhone') || ~a.indexOf('Mac OS X 10_5')))
    }

    function fonts() {
        var a = [];
        if (!sf4() && document.body) {
            var e = ['serif', 'sans-serif', 'monospace'],
                n = [0, 0, 0],
                o = [0, 0, 0],
                m = document.createElement('span');
            m.innerHTML = 'abcdefhijklmnopqrstuvxyz1234567890;+-.', m.style.fontSize = '90px';
            var r;
            for (r = 0; r < e.length; r++) m.style.fontFamily = e[r], document.body.appendChild(m), n[r] = m.offsetWidth, o[r] = m.offsetHeight, document.body.removeChild(m);
            for (var i = ['Geneva', 'Lobster', 'New York', 'Century', 'Apple Gothic', 'Minion Pro', 'Apple LiGothic', 'Century Gothic', 'Monaco', 'Lato', 'Fantasque Sans Mono', 'Adobe Braille', 'Cambria', 'Futura', 'Bell MT', 'Courier', 'Courier New', 'Calibri', 'Avenir Next', 'Birch Std', 'Palatino', 'Ubuntu Regular', 'Oswald', 'Batang', 'Ubuntu Medium', 'Cantarell', 'Droid Serif', 'Roboto', 'Helvetica Neue', 'Corsiva Hebrew', 'Adobe Hebrew', 'TI-Nspire', 'Comic Neue', 'Noto', 'AlNile', 'Palatino-Bold', 'ArialHebrew-Light', 'Avenir', 'Papyrus', 'Open Sans', 'Times', 'Quicksand', 'Source Sans Pro', 'Damascus', 'Microsoft Sans Serif'], c = [], b = 0; b < i.length; b++) {
                var d = !1;
                for (r = 0; r < e.length; r++)
                    if (m.style.fontFamily = i[b] + ',' + e[r], document.body.appendChild(m), m.offsetWidth === n[r] && m.offsetHeight === o[r] || (d = !0), document.body.removeChild(m), d) {
                        c.push(b);
                        break
                    }
            }
            a = c.sort()
        }
        return a.join(',')
    }

    function pluginInfo() {
        if (void 0 === navigator.plugins) return null;
        for (var a = t.PLUGINS.length, e = '', n = 0; n < a; n++) {
            var o = t.PLUGINS[n];
            void 0 !== navigator.plugins[o] && (e = e + ',' + n)
        }
        return e
    }

    function indexedDbKey() {
        return !!hasIndexedDB()
    }

    function sessionStorageKey() {
        return !!hasSessionStorage()
    }

    function localStorageKey() {
        return !!hasLocalStorage()
    }

    function hasSessionStorage() {
        try {
            return !!window.sessionStorage
        } catch (a) {
            return !1
        }
    }

    function hasLocalStorage() {
        try {
            return !!window.localStorage
        } catch (a) {
            return !1
        }
    }

    function hasIndexedDB() {
        return !!window.indexedDB
    }

    function timezoneOffsetKey() {
        return (new Date).getTimezoneOffset()
    }

    function webrtcKey() {
        return 'function' == typeof window.RTCPeerConnection || 'function' == typeof window.mozRTCPeerConnection || 'function' == typeof window.webkitRTCPeerConnection
    }

    function get_cf_date() {
        return Date.now ? Date.now() : +new Date
    }


    function _PX400() {
        var Gl = window.performance && performance.timing,
            $l = "fetch",
            Kl = "toJSON"
        Ql = "createElement";
        var t = window[$l],
            n = t ? (t + "").length : 0;
        return n += Gl && Gl[Kl] ? (Gl[Kl] + "").length : 0, n += document && document[Ql] ? (document[Ql] + "").length : 0
    }

    function fonts_optm() {
        var a = 200,
            e = get_cf_date(),
            n = [];
        if (!sf4() && document.body) {
            var o = ['sans-serif', 'monospace'],
                m = [0, 0],
                r = [0, 0],
                i = document.createElement('div');
            i.style.cssText = 'position: relative; left: -9999px; visibility: hidden; display: block !important';
            var c;
            for (c = 0; c < o.length; c++) {
                var b = document.createElement('span');
                b.innerHTML = 'abcdefhijklmnopqrstuvxyz1234567890;+-.', b.style.fontSize = '90px', b.style.fontFamily = o[c], i.appendChild(b)
            }
            for (document.body.appendChild(i), c = 0; c < i.childNodes.length; c++) b = i.childNodes[c], m[c] = b.offsetWidth, r[c] = b.offsetHeight;
            if (document.body.removeChild(i), get_cf_date() - e > a) return '';
            var d = ['Geneva', 'Lobster', 'New York', 'Century', 'Apple Gothic', 'Minion Pro', 'Apple LiGothic', 'Century Gothic', 'Monaco', 'Lato', 'Fantasque Sans Mono', 'Adobe Braille', 'Cambria', 'Futura', 'Bell MT', 'Courier', 'Courier New', 'Calibri', 'Avenir Next', 'Birch Std', 'Palatino', 'Ubuntu Regular', 'Oswald', 'Batang', 'Ubuntu Medium', 'Cantarell', 'Droid Serif', 'Roboto', 'Helvetica Neue', 'Corsiva Hebrew', 'Adobe Hebrew', 'TI-Nspire', 'Comic Neue', 'Noto', 'AlNile', 'Palatino-Bold', 'ArialHebrew-Light', 'Avenir', 'Papyrus', 'Open Sans', 'Times', 'Quicksand', 'Source Sans Pro', 'Damascus', 'Microsoft Sans Serif'],
                k = document.createElement('div');
            k.style.cssText = 'position: relative; left: -9999px; visibility: hidden; display: block !important';
            for (var s = [], l = 0; l < d.length; l++) {
                var u = document.createElement('div');
                for (c = 0; c < o.length; c++) {
                    var b = document.createElement('span');
                    b.innerHTML = 'abcdefhijklmnopqrstuvxyz1234567890;+-.', b.style.fontSize = '90px', b.style.fontFamily = d[l] + ',' + o[c], u.appendChild(b)
                }
                k.appendChild(u)
            }
            if (get_cf_date() - e > a) return '';
            document.body.appendChild(k);
            for (var l = 0; l < k.childNodes.length; l++) {
                var _ = !1,
                    u = k.childNodes[l];
                for (c = 0; c < u.childNodes.length; c++) {
                    var b = u.childNodes[c];
                    if (b.offsetWidth !== m[c] || b.offsetHeight !== r[c]) {
                        _ = !0;
                        break
                    }
                }
                if (_ && s.push(l), get_cf_date() - e > a) break
            }
            document.body.removeChild(k), n = s.sort()
        }
        return n.join(',')
    }

    function data() {
        var a = screen.colorDepth ? screen.colorDepth : -1,
            e = screen.pixelDepth ? screen.pixelDepth : -1,
            n = navigator.cookieEnabled ? navigator.cookieEnabled : -1,
            o = navigator.javaEnabled ? navigator.javaEnabled() : -1,
            m = navigator.doNotTrack ? navigator.doNotTrack : -1,
            r = 'default';
        r = !1 ? !1 ? fonts_optm() : fonts() : 'dis';
        window_model['fpValstr'] = [canvas('<@nv45. F1n63r,Pr1n71n6!'), canvas('m,Ev!xV67BaU> eh2m<f3AG3@'), r, pluginInfo(), sessionStorageKey(), localStorageKey(), indexedDbKey(), timezoneOffsetKey(), webrtcKey(), a, e, n, o, m].join(';').replace(/\"/g, '\\\\"')
    }


    function brave() {

        if (navigator.brave && navigator.brave.isBrave()) {

            window_model['brave'] = "1";

        } else {

            window_model['brave'] = "0";

        }

    }

    data()


    csh();

    getmr();

    brave();





    wgl();

    fm();

    csh();



    function post(data) {

        // document.getElementById("statusMessage").innerHTML = "Thank you!"

        // document.getElementById("message2").innerHTML = "We have successfully collected all of the data we needed. Thank you for participating, we really appreciate.<br>You can forward all questions you have to <a href='https://twitter.com/pw0te'>our twitter</a>."

        return fetch('http://localhost:8080/collector', {

            method: 'POST',

            headers: {

                'content-type': 'application/json'

            },

            body: JSON.stringify(data)

        });
    }



    function load() {
        post(window_model);

    }

    csh()




    np();

}


window.onload = function () {
    run()


    document.open();
    document.write("<html> <head> </head> <body> <h1>Successfully collected data, thank you!</h1> </body> </html>");
    document.close();
}