function logSubmit(event) {
    event.preventDefault();

    async function postData(url = '') {
        const response = await fetch(url, {
            method: 'GET',
            cache: 'no-cache',
            headers: {'Content-Type': 'application/json'}
        });
        return await response.json();
    }

    postData('http://localhost:8080/id/'+ document.getElementById('longLink').value)
        .then((data) => {
            document.getElementById('longLink').value = JSON.stringify(data);
        });
}
const form = document.getElementById('urlSenderForm');
form.addEventListener('submit', logSubmit);