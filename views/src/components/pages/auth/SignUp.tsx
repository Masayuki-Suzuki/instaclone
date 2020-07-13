import { css, Global } from '@emotion/core'
import React from 'react'
import globalStyle from '~/libs/globalStyle'
import SignUpCard from '~/components/organisms/SignUpCard'
import GetTheAppText from '~/components/atoms/GetTheAppText'
import GetTheAppImage from '~/components/atoms/GetTheAppImage'
// import { RouteComponentProps } from 'react-router-dom'

// type SignUpProps = RouteComponentProps<{}>

const article = css({
    display: 'flex',
    alignItems: 'center',
    flexDirection: 'column',
    justifyContent: 'center',
    height: '100%'
})

const SignUp = (): JSX.Element => {
    return (
        <article css={article}>
            <Global styles={globalStyle} />
            <SignUpCard />
            <div>
                <GetTheAppText />
                <GetTheAppImage />
            </div>
        </article>
    )
}

export default SignUp
