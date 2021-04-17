module.exports = {
  "env": {
    "es6": true,
    "browser": true
  },
  "globals": {
    "Atomics": "readonly",
    "SharedArrayBuffer": "readonly"
  },
  "parser": "babel-eslint",
  "parserOptions": {
    "ecmaVersion": 2018,
    "sourceType": "module",

    "ecmaFeatures": {
      "jsx": true,
      "impliedStrict": true,
      "globalReturn": false,
      "experimentalObjectRestSpread": true,
      "legacyDecorators": true
    }
  },
  "plugins": [
    "react"
  ],
  "extends": [
    "eslint:recommended",
    "plugin:react/recommended"
    //"plugin:prettier/recommended"
  ],
  rules: {
    'no-console':'off',
    /*'indent': [
      'error',
      2
    ],
    'linebreak-style': [
      'error',
      'unix'
    ],
    'quotes': [
      'error',
      'single'
    ],
    'semi': [
      'error',
      'never'
    ],*/
    'react/prop-types': 0,
    'no-mixed-spaces-and-tabs': [2, 'smart-tabs']
  },
};
