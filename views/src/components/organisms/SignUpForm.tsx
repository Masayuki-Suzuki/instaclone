import React, { useState } from 'react'
import { Field, InjectedFormProps, reduxForm } from 'redux-form'
import { css } from '@emotion/core'
import SnsLoginButton from '~/components/atoms/SnsLoginButton'

const container = css({
    padding: '0 40px',
    width: '100%'
})

const form = css({
    display: 'block',
    width: '100%'
})

const inputContainer = css({
    marginBottom: '6px',
    width: '100%'
})

const inputField = css({
    background: '#fafafa',
    border: 'solid 1px #dbdbdb',
    borderRadius: '3px',
    display: 'block',
    fontSize: '1.6rem',
    '-webkit-appearance': 'none',
    padding: '9px 0 7px 8px',
    width: '100%',
    '::placeholder': {
        color: '#aaa',
        fontSize: '12px',
        letterSpacing: '.04em'
    }
})

const login = css({
    padding: '8px 0'
})

const terms = css({
    color: '#8e8e8e',
    fontSize: '1.2rem',
    letterSpacing: '.02em',
    lineHeight: '16px',
    textAlign: 'center',
    padding: '10px 0',
    span: {
        fontWeight: 600
    }
})

const SignUpForm = (props: InjectedFormProps): JSX.Element => {
    const { handleSubmit } = props

    const [validation, setValidation] = useState(false)

    const onSubmitForm = async formProps => {
        console.log(formProps)
    }

    return (
        <div css={container}>
            <form css={form} onSubmit={handleSubmit(onSubmitForm)}>
                <div css={inputContainer}>
                    <Field css={inputField} name="email" type="email" component="input" placeholder="Email" />
                </div>
                <div css={inputContainer}>
                    <Field css={inputField} name="fullName" type="text" component="input" placeholder="Full Name" />
                </div>
                <div css={inputContainer}>
                    <Field css={inputField} name="userName" type="text" component="input" placeholder="Username" />
                </div>
                <div css={inputContainer}>
                    <Field css={inputField} name="password" type="password" component="input" placeholder="Password" />
                </div>
                <div css={login}>
                    <SnsLoginButton text={'Sign up'} validation={validation} />
                </div>
            </form>
            <p css={terms}>
                By signing up, you agree to our <span>Terms</span> , <span>Data</span> <span>Policy</span> and{' '}
                <span>Cookies Policy</span> .
            </p>
        </div>
    )
}

export default reduxForm({ form: 'signUp' })(SignUpForm)
