const Controller = {
    search: (ev) => {
        ev.preventDefault();
        const form = document.getElementById("form");
        const data = Object.fromEntries(new FormData(form));
        const response = fetch(`/search?q=${data.query}`).then((response) => {
            response.json().then((results) => {
                Controller.updateTable(results);
            });
        });
    },

    updateTable: (response) => {
        const tableHead = document.getElementById("table-head");
        const tableBody = document.getElementById("table-body");

        tableHead.innerHTML = `<th>Number of matches : ${response.count}</th>`;

        const rows = [];
        for (let match of response.matches) {
            rows.push(`<td>${match}<tr/>`);
        }
        tableBody.innerHTML = rows;
    },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
