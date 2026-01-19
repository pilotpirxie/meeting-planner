import js from "@eslint/js";
import pluginReact from "eslint-plugin-react";
import pluginReactHooks from "eslint-plugin-react-hooks";
import simpleImportSort from "eslint-plugin-simple-import-sort";
import { defineConfig } from "eslint/config";
import globals from "globals";
import tseslint from "typescript-eslint";

export default defineConfig([
  {
    ignores: ["dist/**", "node_modules/**", "*.d.ts"]
  },
  {
    files: ["**/*.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node
      },
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        project: ["./tsconfig.json", "./tsconfig.app.json", "./tsconfig.node.json"]
      }
    },
    plugins: {
      react: pluginReact,
      "react-hooks": pluginReactHooks,
      "simple-import-sort": simpleImportSort,
    }
  },
  js.configs.recommended,
  ...tseslint.configs.strictTypeChecked,
  ...tseslint.configs.stylisticTypeChecked,
  pluginReact.configs.flat.recommended,
  {
    files: ["**/*.js", "**/*.mjs", "**/*.cjs"],
    ...tseslint.configs.disableTypeChecked
  },
  {
    settings: {
      react: {
        version: "detect"
      }
    },
    rules: {
      quotes: ["error", "double"],
      indent: ["error", 2],
      semi: ["error", "always"],
      "no-console": ["error", { allow: ["warn", "error", "info", "table"] }],
      "no-trailing-spaces": "error",
      eqeqeq: "error",
      "no-mixed-operators": ["error"],

      "react/react-in-jsx-scope": "off",
      "react/jsx-indent-props": ["error", 2],
      "react/boolean-prop-naming": [
        "error",
        { rule: "^(is|are|has|should|can)[A-Z]([A-Za-z0-9]?)+" },
      ],
      "react/jsx-curly-brace-presence": ["error", { props: "never", children: "ignore" }],
      "react/no-array-index-key": "error",
      "react/jsx-no-leaked-render": ["error", { validStrategies: ["ternary"] }],
      "react/no-unused-prop-types": "error",
      "react/jsx-no-useless-fragment": "error",
      "react/self-closing-comp": "error",
      "react/jsx-first-prop-new-line": ["error", "multiline"],
      "react/jsx-closing-bracket-location": [
        "error",
        { selfClosing: "line-aligned", nonEmpty: "after-props" },
      ],
      "react/jsx-max-props-per-line": ["error", { maximum: 1 }],

      "react/prop-types": "off",
      "react-hooks/rules-of-hooks": "error",
      "react-hooks/exhaustive-deps": "warn",
      "@typescript-eslint/no-unused-vars": [
        "error",
        {
          argsIgnorePattern: "^_",
          varsIgnorePattern: "^_"
        }
      ],
      "@typescript-eslint/prefer-nullish-coalescing": "off",
    }
  }
]);
