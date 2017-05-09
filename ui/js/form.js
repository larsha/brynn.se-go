class Form {
  constructor (form) {
    this.form = form;
    this.modal = document.querySelector('#form-confirmation')
    this.modalMessage = document.querySelector('#form-confirmation-message')
    this.body = document.querySelector('body')
    this.emailElement = this.form.querySelector('#email')
    this.messageElement = this.form.querySelector('#message')
    this.successClass = 'form-confirmation--success'
    this.errorClass = 'form-confirmation--error'
    this.modalClass = 'form-confirmation-open'
    this.events()
  }

  get email () {
    return this.emailElement.value
  }

  get message () {
    return this.messageElement.value
  }

  events () {
    this.form.addEventListener('submit', e => this.submit(e))
  }

  setModal (className, message = '') {
    this.body.classList.add(this.modalClass)
    this.modal.classList.add(className)
    this.modalMessage.innerHTML = message

    setTimeout(() => {
      this.body.classList.remove(this.modalClass)
    }, 3000)
  }

  submit (e) {
    e.stopPropagation()
    e.preventDefault()
    this.modal.classList.remove(this.successClass)
    this.modal.classList.remove(this.errorClass)

    const options = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        email: this.email,
        message: this.message
      })
    }

    fetch('/api/form', options)
      .then(res => {
        if (res.status >= 400 && res.status <= 599) {
          return Promise.reject()
        }

        this.emailElement.value = ''
        this.messageElement.value= ''
        this.setModal(this.successClass, 'Thank you!')
      })
      .catch(() => {
        this.emailElement.focus()
        this.setModal(this.errorClass, 'Sorry!')
      });
  }
}

export default (() => {
  const form = document.querySelector('#form')
  if (form.length > 0) {
    new Form(form);
  }
})()
