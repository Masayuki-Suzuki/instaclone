import React from 'react'
import { css } from '@emotion/core'
import { ReactComponent as LogoSvg } from '../../assets/images/logo.svg'

const logo = css({
    display: 'block',
    maxHeight: '50px',
    margin: '22px auto 18px'
})

const Logo = (): JSX.Element => {
    return <LogoSvg css={logo} />
}

export default Logo
