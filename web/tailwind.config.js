const theme = {
    extend: {
        backgroundImage: {
            "background-texture": "url('/assets/background-texture.png')",
        },
        fontFamily: {
            sans: ["Jost", "sans-serif"],
        },
    },
};

module.exports = {
    content: ["components/**/*.templ"],
    theme,
    plugins: [],
};
