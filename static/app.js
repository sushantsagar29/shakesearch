const form = document.getElementById("form");
const tableHead = document.getElementById("table-head");
const tableBody = document.getElementById("table-body");
const loader = document.getElementById("loader-wrapper");

const Controller = {
    search: (ev) => {
        ev.preventDefault();
        const data = Object.fromEntries(new FormData(form));
        loader.style.display = "flex";
        const response = fetch(`/search?q=${data.query}&sensitive=${data.case !== undefined}&exact=${data.word !== undefined}`)
            .then((response) => {
                if (response.status != 200) {
                    Controller.handleError()
                }
                response.json().then((results) => {
                    Controller.updateTable(results);
                });
            }).catch(() => {
                Controller.handleError()
            });
    },

    updateTable: (response) => {
        tableHead.innerHTML = `<th colspan="2">Number of matches : ${response.count}</th>`;

        let content = "";
        for (let i = 0; i < response.matches.length; i++) {
            content += (`<tr>`);
            content += (`<td>${i + 1}</td>`);
            content += (`<td>${response.matches[i]}</td>`);
            content += (`</tr>`);
        }

        loader.style.display = "none";
        tableBody.innerHTML = content;
    },

    handleError: () => {
        tableHead.innerHTML = `<th colspan="2">Oops! Something went wrong.</th>`;
        loader.style.display = "none";
    },
};

form.addEventListener("submit", Controller.search);
