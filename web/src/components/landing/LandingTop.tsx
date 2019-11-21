import React from "react"
import { landingPageStyles } from "../../styles/pageStyles"

// baseweb
import { Block } from "baseui/block"
import { Display1 } from "baseui/typography"
import { Button } from "baseui/button"

function LandingTop() {
	return (
		<div>
			<Block
				overrides={{
					Block: {
						style: landingPageStyles.landing,
					},
				}}
			>
				<Block
					overrides={{
						Block: {
							style: { ...landingPageStyles.welcomeContent, textAlign: "center" },
						},
					}}
				>
					<Display1 overrides={{ Block: { style: landingPageStyles.welcomeHeader } }}>Welcome to GyfHUB</Display1>
					<Button
						overrides={{
							BaseButton: {
								style: landingPageStyles.joinBtn,
							},
						}}
					>
						Join!
					</Button>
				</Block>
			</Block>
		</div>
	)
}

export default LandingTop
