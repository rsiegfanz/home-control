/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
    content: ['./src/**/*.{html,ts,tsx,js,jsx}'],
    theme: {
        extend: {
            colors: {
                primary: 'var(--color-primary)',
                secondary: 'var(--color-secondary)',
                success: 'var(--color-success)',
                danger: 'var(--color-danger)',
                info: 'var(--color-info)',
                label: 'var(--color-label)',
                page: 'var(--color-page)',
            },
            fontSize: {
                lg: 'var(--font-primary)',
                md: 'var(--font-label)',
            },
            fontFamily: {
                serif: ['var(--font-family-serif)', ...defaultTheme.fontFamily.serif],
                sans: ['var(--font-family-sans)', ...defaultTheme.fontFamily.sans],
                mono: ['var(--font-family-mono)', ...defaultTheme.fontFamily.sans],
            },
        },
    },
};
