import React from 'react'
import { Link } from 'react-router-dom'
import { css } from '@emotion/core'

type LeadToSignInProps = {
    signUp?: boolean
}

const signUpLink = css({
    color: '#262626',
    fontSize: '1.4rem',
    textAlign: 'center',
    padding: '15px 0',
    a: {
        color: '#0095f6',
        textDecoration: 'none'
    }
})

const link = css({
    fontWeight: 600
})

const LeadToSignInUp = ({ signUp }: LeadToSignInProps): JSX.Element => {
    if (signUp) {
        return (
            <p css={signUpLink}>
                Don't have an account?{' '}
                <Link css={link} to="/accounts/signup">
                    Sign Up
                </Link>
            </p>
        )
    }
    return (
        <p css={signUpLink}>
            Have an account? <Link to="/accounts/signin">Log in</Link>
        </p>
    )
}

export default LeadToSignInUp
