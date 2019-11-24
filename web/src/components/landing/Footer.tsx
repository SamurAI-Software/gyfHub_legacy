import React from "react"
import { FlexGrid, FlexGridItem } from "baseui/flex-grid"
import { BlockProps } from "baseui/block"

interface Props {}
const itemProps: BlockProps = {
	backgroundColor: "transparent",
	display: "flex",
	alignItems: "center",
	justifyContent: "center",
	color: "#fff",
}

const bottomFooter: BlockProps = {
	display: "flex",
	// alignItems: "center",
	justifyContent: "left",
	color: "#fff",
	marginBottom: "4rem",
	padding: "3rem",
}

export default function Footer(props: Props) {
	const {} = props

	return (
		<div className="Footer">
			<div className="Footer-content">
				<ul>
					<h1 className="Footer-title">NAVIGATION</h1>
					<li>
						<a href="http://" target="_blank" rel="noopener noreferrer">
							Back to Top
						</a>
					</li>
					<li>
						<a href="http://" target="_blank" rel="noopener noreferrer">
							What is GyfHUB
						</a>
					</li>
					<li>
						<a href="http://" target="_blank" rel="noopener noreferrer">
							Video
						</a>
					</li>
				</ul>
				<ul>
					<h1 className="Footer-title">THE TEAM</h1>

					<li>
						<a href="http://" target="_blank" rel="noopener noreferrer">
							Meet the team
						</a>
					</li>
				</ul>
				<ul>
					<li>
						<img
							src="https://www.ninjasoftware.com.au/wp-content/themes/ninjadojo/images/animation/NinjaSoftwareLogoLogomark.svg"
							alt=""
							className="compLogo"
						/>
					</li>
				</ul>
			</div>
		</div>
	)
}

// export default Footer
