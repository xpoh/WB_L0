function logSubmit(event) {
    event.preventDefault();

    async function postData(url = '') {
        const response = await fetch(url, {
            method: 'GET',
            cache: 'no-cache',
        });
        return await response.json();
    }

    postData('http://localhost/id/'+ document.getElementById('longLink').value)
        .then((data) => {
            document.getElementById('longLink').value = data.url;
        });
}
const form = document.getElementById('urlSenderForm');
form.addEventListener('submit', logSubmit);