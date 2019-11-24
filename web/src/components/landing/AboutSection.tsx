import React from "react"
import { Block } from "baseui/block"
import { landingPageStyles } from "../../styles/pageStyles"
import { Display1 } from "baseui/typography"

function About() {
	return (
		<div>
			<Block
				overrides={{
					Block: {
						style: { ...landingPageStyles.aboutSection, textAlign: "center" },
					},
				}}
			>
				<Block
					overrides={{
						Block: {
							style: {},
						},
					}}
				>
					<Display1 overrides={{ Block: { style: landingPageStyles.sectionHeader } }}>What is GyfHUB</Display1>
					<Block overrides={{ Block: { style: { ...landingPageStyles.aboutContent, textAlign: "left" } } }}>
						<div>
							<p>
								Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum impedit reiciendis cumque quia numquam laudantium ipsa minima magni similique
								veritatis placeat, tenetur aliquid a odit esse blanditiis vero, maiores, deserunt soluta reprehenderit. Illum nisi praesentium omnis eligendi!
								Ipsa ratione vel laborum rerum numquam quas debitis, voluptates veniam quidem, qui ipsdeserunt soluta reprehenderit. Illum nisi praesentium
								omnis eligendi! Ipsa ratione vel laborum rerum numquam quas debitis, voluptates Lorem ipsum dolor sit amet consectetur adipisicing elit.
								Expedita suscipit quia eaque accusamus libero. Impedit itaque, voluptatibus reiciendis facere quas possimus porro. veniam quidem, qui ipsum
								placeat deserunt itaqu
							</p>

							<p>
								Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam omnis quisquam, quis quam tempore aliquam quo illo animi doloribus, at
								praesentium commodi magni porro debitis corporis obcaecati ullam rerum maxime temporibus officiis modi suscipit. Reprehenderit laborum qui
								soluta autem quisquam, provident sapiente asperiores temporibus!
							</p>
						</div>

						<Block>
							<Block overrides={{ Block: { style: landingPageStyles.aboutImg } }}></Block>
						</Block>
					</Block>
				</Block>
			</Block>
		</div>
	)
}

export default About
