import React from 'react'
import { css } from '@emotion/core'
import Logo from '~/components/atoms/Logo'
import SignUpCaption from '~/components/atoms/SignUpCaption'
import SnsLoginButton from '~/components/atoms/SnsLoginButton'

const snsLogin = css({
    padding: '6px 40px'
})

const SignUpHeader = (): JSX.Element => {
    return (
        <header>
            <Logo />
            <SignUpCaption />
            <div css={snsLogin}>
                <SnsLoginButton text="Log in with Google" brand="google" validation={true} />
            </div>
        </header>
    )
}

export default SignUpHeader
