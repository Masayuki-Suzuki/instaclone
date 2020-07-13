import React from 'react'
import { css } from '@emotion/core'

const container = css({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: '10px 0'
})

const wrapper = css({
    marginRight: 8
})

const image = css({
    height: '40px'
})

const GetTheAppImage = (): JSX.Element => {
    return (
        <div css={container}>
            <span css={wrapper}>
                <img css={image} src="/images/appstore.png" alt="App Store" />
            </span>
            <span>
                <img css={image} src="/images/googleplay.png" alt="Goole Play" />
            </span>
        </div>
    )
}

export default GetTheAppImage
