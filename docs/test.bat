@echo off
chcp 65001
setlocal enabledelayedexpansion
cd %~dp0

set ROOT=elasticsearch

rem make root dir named elasticsearch, then cd
mkdir %ROOT%
cd %ROOT%

(
echo {
echo   "systemDict": "system_core.dic",
echo   "inputTextPlugin": [
echo     { "class": "com.worksap.nlp.sudachi.DefaultInputTextPlugin" },
echo     {
echo       "class": "com.worksap.nlp.sudachi.ProlongedSoundMarkInputTextPlugin",
echo       "prolongedSoundMarks": ["ー", "-", "⁓", "〜", "〰"],
echo       "replacementSymbol": "ー"
echo     }
echo   ],
echo   "oovProviderPlugin": [
echo     { "class": "com.worksap.nlp.sudachi.MeCabOovProviderPlugin" },
echo     {
echo       "class": "com.worksap.nlp.sudachi.SimpleOovProviderPlugin",
echo       "oovPOS": ["補助記号", "一般", "*", "*", "*", "*"],
echo       "leftId": 5968,
echo       "rightId": 5968,
echo       "cost": 3857
echo     }
echo   ],
echo   "pathRewritePlugin": [
echo     { "class": "com.worksap.nlp.sudachi.JoinNumericPlugin", "joinKanjiNumeric": true },
echo     {
echo       "class": "com.worksap.nlp.sudachi.JoinKatakanaOovPlugin",
echo       "oovPOS": ["名詞", "普通名詞", "一般", "*", "*", "*"],
echo       "minLength": 3
echo     }
echo   ]
echo }
)>"sudachi.json"
