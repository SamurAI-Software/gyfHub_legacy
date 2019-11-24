import React from "react"
import { landingPageStyles } from "../../styles/pageStyles"

// imgs
import phomeImg from "../../imgs/phoneMockup.png"

// baseweb
import { Block } from "baseui/block"
import { Display1, Display2, Display3, Display4 } from "baseui/typography"
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
				<div style={{ maxWidth: "1300px", display: "flex", justifyContent: "center", margin: "auto" }}>
					<Block
						overrides={{
							Block: {
								style: { ...landingPageStyles.welcomeContent, textAlign: "left" },
							},
						}}
					>
						<div>
							<Display1 overrides={{ Block: { style: landingPageStyles.welcomeHeader } }}>Welcome to GyfHUB</Display1>
							<Display4 overrides={{ Block: { style: landingPageStyles.welcomeHeader } }}>Messaging platform for GIF lovers.</Display4>
						</div>
						<Button
							overrides={{
								BaseButton: {
									style: { ...landingPageStyles.joinBtn, justifyContent: "center" },
								},
							}}
						>
							Join!
						</Button>
					</Block>

					<Block
						overrides={{
							Block: {
								style: { ...landingPageStyles.welcomeContent, textAlign: "center" },
							},
						}}
					>
						<img src={phomeImg} style={landingPageStyles.phoneImg} alt="" />
					</Block>
				</div>
			</Block>
		</div>
	)
}

export default LandingTop
