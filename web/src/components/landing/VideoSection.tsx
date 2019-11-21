import React from "react"
import { Block } from "baseui/block"
import { landingPageStyles } from "../../styles/pageStyles"
import { Display1 } from "baseui/typography"

function VideoSection() {
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
					<Display1 overrides={{ Block: { style: landingPageStyles.sectionHeader } }}>Video</Display1>
					<Block overrides={{ Block: { style: { ...landingPageStyles.videoContent, textAlign: "left" } } }}>
						<Block>
							<Block overrides={{ Block: { style: landingPageStyles.videoImg } }}>
								<iframe width="420" title="video" height="315" src="https://www.youtube.com/embed/ycPr5-27vSI"></iframe>
							</Block>
						</Block>
						<div>
							<p>
								Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum impedit reiciendis cumque quia numquam laudantium ipsa minima magni similique
								veritatis placeat, tenetur aliquid a odit esse blanditiis vero, maiores, deserunt soluta reprehenderit. Illum nisi praesentium omnis eligendi!
								Ipsa ratione vel laborum rerum numquam quas debitis, voluptates veniam quidem, qui ipsdeserunt soluta reprehenderit. Illum nisi praesentium
								omnis eligendi! Ipsa ratione vel laborum rerum numquam quas debitis, voluptates Lorem ipsum dolor sit amet consectetur adipisicing elit.
								Expedita suscipit quia eaque accusamus libero. Impedit itaque, voluptatibus reiciendis facere quas possimus porro. veniam quidem, qui ipsum
								placeat deserunt itaque error fugit? Lorem ipsum dolor sit amet consectetur adipisicing elit. Cumque minus blanditiis at consequatur? Sit beatae
								placeat illo iusto nobis, labore iste assumenda voluptas quia ad perspiciatis quasi ipsam rem eos suscipit odio numquam quis quaerat sed. Fugit
								ea doloribus perspiciatis.
							</p>
							<p>
								Lorem ipsum dolor sit amet, consectetur adipisicing elit. Reprehenderit fuga dolorum qui iusto, iure deserunt earum minus tempore, doloremque
								quibusdam sit molestias nam nisi recusandae inventore, doloribus enim voluptates! Sed!
							</p>
							<p>
								Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsam omnis quisquam, quis quam tempore aliquam quo illo animi doloribus, at
								praesentium commodi magni porro debitis corporis obcaecati ullam rerum maxime temporibus officiis modi suscipit. Reprehenderit laborum qui
								soluta autem quisquam, provident sapiente asperiores temporibus!
							</p>
						</div>
					</Block>
				</Block>
			</Block>
		</div>
	)
}

export default VideoSection
