const Regex = {
    email: /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/,
    password: /^(?=.*?[a-zA-Z])(?=.*?\d)[!-\~]{8,255}$/
}

export default {
    empty: value => (value || typeof value === 'number' ? undefined : 'Required'),
    email: value => (value && !Regex.email.test(value) ? 'error' : undefined),
    password: value => (value && !Regex.password.test(value) ? 'error' : undefined)
}
