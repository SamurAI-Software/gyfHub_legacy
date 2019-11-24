import React from "react"

// styles
import { global } from "../../styles/pageStyles"

interface Props {}

function GifGrid(props: Props) {
	const {} = props

	return (
		<global.CenteredNav>
			<div className="GifGrid">
				<img src="https://media3.giphy.com/media/3oEhmLyf84WAWAEnRe/giphy.gif" alt="" />
				<img src="https://media2.giphy.com/media/26gJyIscAHtBNcc00/giphy.gif" alt="" />
				<img src="https://thumbs.gfycat.com/FeistySelfreliantCrayfish-small.gif" alt="" />
				{/* <img src="https://media0.giphy.com/media/Lcn0yF1RcLANG/source.gif" alt="" /> */}
				{/* <img src="https://media2.giphy.com/media/4cuyucPeVWbNS/giphy.gif" alt="" /> */}
				<img src="https://media1.giphy.com/media/2oPODTJEUNeJG/giphy.gif" alt="" />
				<img src="https://thumbs.gfycat.com/AntiqueRemoteFlounder-max-1mb.gif" alt="" />
				<img src="https://media.giphy.com/media/t8L4eNho9MyOyMNTJc/giphy.gif" alt="" />
			</div>
		</global.CenteredNav>
	)
}

export default GifGrid
