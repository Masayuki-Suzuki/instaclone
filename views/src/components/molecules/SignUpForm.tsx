import React, { useEffect } from 'react'
import { Field, InjectedFormProps, reduxForm } from 'redux-form'
import { css } from '@emotion/core'
import ReduxFormInputIField from './ReduxFormInputField'
import Firebase from '~/libs/Firebase'
import SnsLoginButton from '~/components/atoms/SnsLoginButton'
import validations from '~/libs/validations'
import SignUpTerms from '~/components/atoms/SignUpTerms'

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

const login = css({
    padding: '8px 0'
})

const SignUpForm = (props: InjectedFormProps): JSX.Element => {
    const { handleSubmit, valid } = props

    useEffect(() => {
        Firebase.auth().onAuthStateChanged(user => {
            console.log(user)
            if (user) {
                console.log('Has already logged in.')
                console.log(user)
                // ToDo: be coding redirect functionality to explorer page.
            }
        })
    }, [])

    const onSubmitForm = async formProps => {
        console.log(formProps)
    }

    return (
        <div css={container}>
            <form css={form} onSubmit={handleSubmit(onSubmitForm)}>
                <div css={inputContainer}>
                    <Field
                        name="email"
                        type="email"
                        component={ReduxFormInputIField}
                        placeholder="Email"
                        validate={[validations.empty, validations.email]}
                    />
                </div>
                <div css={inputContainer}>
                    <Field
                        name="fullName"
                        type="text"
                        component={ReduxFormInputIField}
                        placeholder="Full Name"
                        validate={[validations.empty]}
                    />
                </div>
                <div css={inputContainer}>
                    <Field
                        name="userName"
                        type="text"
                        component={ReduxFormInputIField}
                        placeholder="Username"
                        validate={[validations.empty]}
                    />
                </div>
                <div css={inputContainer}>
                    <Field
                        name="password"
                        type="password"
                        component={ReduxFormInputIField}
                        placeholder="Passwordr"
                        validate={[validations.empty]}
                    />
                </div>
                <div css={login}>
                    <SnsLoginButton text={'Sign up'} validation={valid} />
                </div>
            </form>
            <SignUpTerms />
        </div>
    )
}

export default reduxForm({ form: 'signUp' })(SignUpForm)
