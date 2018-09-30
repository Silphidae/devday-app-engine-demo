window.onload = () => {
    getCandy();
    document.getElementById("candy-form").onsubmit = (e) => handleSubmit(e);
};

const getCandy = () => {
    fetch('/api', { headers: { 'Content-Type': 'application/json' } })
    .then(response => response.json())
    .then((results) => {
        if (results.items) {
            results.items.map((item) => {
                addToList(item.Text);
            })
        }
    });
};

const handleSubmit = (e) => {
    const textInput = document.getElementById("candy-text");
    if (textInput.value.length > 0) {
        // addToList(textInput.value);
        save(textInput.value);
    }
    e.preventDefault();
};

const save = (value) => {
    fetch('/api', {
        method: 'POST',
        body: JSON.stringify({ text: value }),
        headers: { 'Content-Type': 'application/json' }
    })
    .then(response => response.json())
    .then((result) => {
        addToList(result.Text);
    });
}

const addToList = (val) => {
    const parent = document.getElementById("candy-list");
    const el = document.createElement('li');
    el.appendChild(document.createTextNode(val));
    el.style.backgroundColor =  `hsla(${Math.random() * 360}, 100%, 70%, 1)`;
    parent.appendChild(el);
    document.getElementById("candy-text").value = "";
};
