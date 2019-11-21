import React from "react"

import { HeaderNavigation, ALIGN, StyledNavigationList, StyledNavigationItem } from "baseui/header-navigation"
import { StyledLink } from "baseui/link"
import { Button } from "baseui/button"

function Navbar() {
	return (
		<HeaderNavigation overrides={{ Root: { style: { position: "fixed", top: 0, minWidth: "100vw", background: "#fff" } } }}>
			<StyledNavigationList $align={ALIGN.left}>
				<StyledNavigationItem>GyfHUB</StyledNavigationItem>
			</StyledNavigationList>
			<StyledNavigationList $align={ALIGN.center} />
			<StyledNavigationList $align={ALIGN.right}>
				<StyledNavigationItem>
					<StyledLink href="#basic-link1">Tab Link One</StyledLink>
				</StyledNavigationItem>
				<StyledNavigationItem>
					<StyledLink href="#basic-link2">Tab Link Two</StyledLink>
				</StyledNavigationItem>
			</StyledNavigationList>
			<StyledNavigationList $align={ALIGN.right}>
				<StyledNavigationItem>
					<Button>Get started</Button>
				</StyledNavigationItem>

				<StyledNavigationItem></StyledNavigationItem>
			</StyledNavigationList>
		</HeaderNavigation>
	)
}

export default Navbar
