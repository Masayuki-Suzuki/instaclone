import React from 'react'
import { css } from '@emotion/core'

const caption = css({
    color: '#9e9e9e',
    fontSize: '1.7rem',
    fontWeight: 500,
    letterSpacing: '.02em',
    lineHeight: '20px',
    textShadow: '0 0 .5px #9e9e9e',
    margin: '0 32px 14px',
    textAlign: 'center'
})

const SignUpCaption = (): JSX.Element => {
    return <h2 css={caption}>Sign up to see photos and videos from your friends.</h2>
}

export default SignUpCaption
