import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';

export const MONKEY_LANGUAGE = "monkey"
export const MonarchConfig: Monaco.languages.IMonarchLanguage = {

    keywords: [
        'let', 'true', 'false', 'if', 'else', 'return', 'fn'
    ],

    operators: [
        '=', '>', '<', '!', '==', '!=', "*", "/", "+", "-"
    ],
    // The main tokenizer for our languages
    tokenizer: {
        root: [
            // identifiers and keywords
            [/[a-z_$][\w$]*/, {
                cases: {
                    '@keywords': 'keyword',
                    '@default': 'identifier'
                }
            }],

        ]
    },
};

export const MonacoLanguageConfiguration: Monaco.languages.LanguageConfiguration = {
    surroundingPairs: [
        { open: '{', close: '}' },
        { open: '[', close: ']' },
        { open: '(', close: ')' },
        { open: '"', close: '"' },
    ],
    autoClosingPairs: [
        { open: '{', close: '}' },
        { open: '[', close: ']' },
        { open: '(', close: ')' },
        { open: '"', close: '"', notIn: ['string', 'comment'] },
    ],

}