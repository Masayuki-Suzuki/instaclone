import React from 'react'
import { css } from '@emotion/core'
import { connect } from 'react-redux'
import { googleLogin } from '~/actions'
import Logo from '~/components/atoms/Logo'
import SignUpCaption from '~/components/atoms/SignUpCaption'
import SnsLoginButton from '~/components/atoms/SnsLoginButton'

type SignUpHeaderProps = {
    googleLogin: (isSignedIn: boolean) => void
}

const snsLogin = css({
    padding: '6px 40px'
})

const SignUpHeader = ({ googleLogin }: SignUpHeaderProps): JSX.Element => {
    // ToDo: will get isSignedIn state from redux for googleLogin's argument.
    const onGoogleLogin = () => googleLogin(false)

    return (
        <header>
            <Logo />
            <SignUpCaption />
            <div css={snsLogin}>
                <SnsLoginButton
                    text="Log in with Google"
                    brand="google"
                    validation={true}
                    onClickButton={onGoogleLogin}
                />
            </div>
        </header>
    )
}

export default connect(null, { googleLogin })(SignUpHeader)
