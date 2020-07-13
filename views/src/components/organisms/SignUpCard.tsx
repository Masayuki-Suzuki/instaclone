import React from 'react'
import { css } from '@emotion/core'
import SignUpHeader from '~/components/molecules/SignUpHeader'
import Separator from '~/components/atoms/Separator'
import SignUpForm from '~/components/organisms/SignUpForm'
import LeadToSignInUp from '~/components/atoms/LeadToSignInUp'

const card = css({
    background: '#fff',
    border: 'solid 1px #ddd',
    borderRadius: '1px',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    marginBottom: '10px',
    maxWidth: '350px',
    width: '100%',
    padding: '10px 0'
})

const SignUpCard = (): JSX.Element => {
    return (
        <React.Fragment>
            <div css={card}>
                <SignUpHeader />
                <Separator />
                <SignUpForm />
            </div>
            <div css={card}>
                <LeadToSignInUp />
            </div>
        </React.Fragment>
    )
}

export default SignUpCard
