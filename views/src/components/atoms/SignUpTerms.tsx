import React from 'react'
import { css } from '@emotion/core'

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

const SignUpTerms = (): JSX.Element => {
    return (
        <p css={terms}>
            By signing up, you agree to our <span>Terms</span> , <span>Data</span> <span>Policy</span> and{' '}
            <span>Cookies Policy</span> .
        </p>
    )
}

export default SignUpTerms
