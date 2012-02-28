// THIS FILE IS AUTO-GENERATED FROM SHA1.js
// DO NOT EDIT.
package ui
func init() {
	Files.Add("SHA1.js", "// From http://code.google.com/p/crypto-js/\r\n// License: http://www.opensource.org/licenses/bsd-license.php\r\n//\r\n// Copyright (c) 2009, Jeff Mott. All rights reserved.\r\n// \r\n// Redistribution and use in source and binary forms, with or without\r\n// modification, are permitted provided that the following conditions are met:\r\n// \r\n// Redistributions of source code must retain the above copyright notice, this\r\n// list of conditions and the following disclaimer. Redistributions in binary\r\n// form must reproduce the above copyright notice, this list of conditions and\r\n// the following disclaimer in the documentation and/or other materials provided\r\n// with the distribution. Neither the name Crypto-JS nor the names of its\r\n// contributors may be used to endorse or promote products derived from this\r\n// software without specific prior written permission. THIS SOFTWARE IS PROVIDED\r\n// BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND ANY EXPRESS OR IMPLIED\r\n// WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF\r\n// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO\r\n// EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,\r\n// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES\r\n// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;\r\n// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND\r\n// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT\r\n// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS\r\n// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\r\n\r\n(function(){\r\n\r\n// Shortcuts\r\nvar C = Crypto,\r\n    util = C.util,\r\n    charenc = C.charenc,\r\n    UTF8 = charenc.UTF8,\r\n    Binary = charenc.Binary;\r\n\r\n// Public API\r\nvar SHA1 = C.SHA1 = function (message, options) {\r\n\tvar digestbytes = util.wordsToBytes(SHA1._sha1(message));\r\n\treturn options && options.asBytes ? digestbytes :\r\n\t       options && options.asString ? Binary.bytesToString(digestbytes) :\r\n\t       util.bytesToHex(digestbytes);\r\n};\r\n\r\n// The core\r\nSHA1._sha1 = function (message) {\r\n\r\n\t// Convert to byte array\r\n\tif (message.constructor == String) message = UTF8.stringToBytes(message);\r\n\t/* else, assume byte array already */\r\n\r\n\tvar m  = util.bytesToWords(message),\r\n\t    l  = message.length * 8,\r\n\t    w  =  [],\r\n\t    H0 =  1732584193,\r\n\t    H1 = -271733879,\r\n\t    H2 = -1732584194,\r\n\t    H3 =  271733878,\r\n\t    H4 = -1009589776;\r\n\r\n\t// Padding\r\n\tm[l >> 5] |= 0x80 << (24 - l % 32);\r\n\tm[((l + 64 >>> 9) << 4) + 15] = l;\r\n\r\n\tfor (var i = 0; i < m.length; i += 16) {\r\n\r\n\t\tvar a = H0,\r\n\t\t    b = H1,\r\n\t\t    c = H2,\r\n\t\t    d = H3,\r\n\t\t    e = H4;\r\n\r\n\t\tfor (var j = 0; j < 80; j++) {\r\n\r\n\t\t\tif (j < 16) w[j] = m[i + j];\r\n\t\t\telse {\r\n\t\t\t\tvar n = w[j-3] ^ w[j-8] ^ w[j-14] ^ w[j-16];\r\n\t\t\t\tw[j] = (n << 1) | (n >>> 31);\r\n\t\t\t}\r\n\r\n\t\t\tvar t = ((H0 << 5) | (H0 >>> 27)) + H4 + (w[j] >>> 0) + (\r\n\t\t\t         j < 20 ? (H1 & H2 | ~H1 & H3) + 1518500249 :\r\n\t\t\t         j < 40 ? (H1 ^ H2 ^ H3) + 1859775393 :\r\n\t\t\t         j < 60 ? (H1 & H2 | H1 & H3 | H2 & H3) - 1894007588 :\r\n\t\t\t                  (H1 ^ H2 ^ H3) - 899497514);\r\n\r\n\t\t\tH4 =  H3;\r\n\t\t\tH3 =  H2;\r\n\t\t\tH2 = (H1 << 30) | (H1 >>> 2);\r\n\t\t\tH1 =  H0;\r\n\t\t\tH0 =  t;\r\n\r\n\t\t}\r\n\r\n\t\tH0 += a;\r\n\t\tH1 += b;\r\n\t\tH2 += c;\r\n\t\tH3 += d;\r\n\t\tH4 += e;\r\n\r\n\t}\r\n\r\n\treturn [H0, H1, H2, H3, H4];\r\n\r\n};\r\n\r\n// Package private blocksize\r\nSHA1._blocksize = 16;\r\n\r\n})();\r\n");
}