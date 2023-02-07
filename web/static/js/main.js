const send_form = document.getElementById("sender")
const status = document.getElementById("status")
const send_btn = document.getElementById("sen-btn")

async function postFormDataAsJson({ url, formData }) {
    const plainFormData = Object.fromEntries(formData.entries());
    const formDataJsonString = JSON.stringify(plainFormData);

    const fetchOptions = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: formDataJsonString,
    };

    const response = await fetch(url, fetchOptions);
    return response.json();
}

async function handleFormSubmit(event) {
    event.preventDefault();

    const form = event.currentTarget;
    const url = form.action;

    try {
        status.innerHTML = "Ожидайте"
        status.style.backgroundColor = "yellow"
        send_btn.disabled = true
        const formData = new FormData(form);
        const responseData = await postFormDataAsJson({ url, formData });
        send_btn.disabled = false
        if (responseData["comment"] !== "OK") {
            status.innerHTML = responseData["comment"]
            status.style.backgroundColor = "red"
        } else {
            status.innerHTML = "OK"
            status.style.backgroundColor = "green"
            window.location.href="/download"
        }
    } catch (error) {
        console.error(error);
    }
}

send_form.addEventListener("submit", handleFormSubmit);