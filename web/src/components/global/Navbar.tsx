import React from "react"

import { global } from "../../styles/pageStyles"

import { HeaderNavigation, ALIGN, StyledNavigationList, StyledNavigationItem } from "baseui/header-navigation"
import { StyledLink } from "baseui/link"
import { Button } from "baseui/button"

function Navbar() {
	return (
		<HeaderNavigation overrides={{ Root: { style: { position: "fixed", top: 0, minWidth: "100vw", background: "#fff", backgroundColor: "#000" } } }}>
			<global.CenteredNav>
				<StyledNavigationList $align={ALIGN.left}>
					<StyledNavigationItem>GyfHUB</StyledNavigationItem>
				</StyledNavigationList>
				<StyledNavigationList $align={ALIGN.center} />
				<StyledNavigationList $align={ALIGN.right}>
					<StyledNavigationItem>
						<StyledLink href="#basic-link1"></StyledLink>
					</StyledNavigationItem>
					<StyledNavigationItem>
						<StyledLink href="#basic-link2"></StyledLink>
					</StyledNavigationItem>
				</StyledNavigationList>
				<StyledNavigationList $align={ALIGN.right}>
					<StyledNavigationItem>
						<Button>SIGN UP!</Button>
					</StyledNavigationItem>

					<StyledNavigationItem></StyledNavigationItem>
				</StyledNavigationList>
			</global.CenteredNav>
		</HeaderNavigation>
	)
}

export default Navbar
