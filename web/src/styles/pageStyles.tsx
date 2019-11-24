import bg from "../imgs/bg1.svg"
import { styled } from "baseui"

// LANDING PAGE
export const global = {
	// sticky: {
	// 	position: "fixed",
	// 	minWidth: "100vw",
	// 	background: "#000",
	// },
	Centered: styled("div", {
		display: "flex",
		overflowX: "hidden",
		flexDirection: "column",
		justifyContent: "center",
		alignItems: "center",
		height: "100%",
	}),

	CenteredNav: styled("div", {
		display: "flex",
		minWidth: "80%",
		justifyContent: "center",
		alignItems: "center",
		margin: "auto",
		// height: "100%",
		maxWidth: "1300px",
	}),
}
export const landingPageStyles = {
	landing: {
		minHeight: "100vh",
		// background: "linear-gradient(90deg, rgba(0,0,0,1) 0%, rgba(48,48,48,1) 46%, rgba(75,75,75,0.9895308465182948) 100%);",
		minWidth: "100vw",
		backgroundImage: `url(${bg})`,
		backgroundAttachment: "fixed",
		backgroundSize: "contain",
		// display: "flex",
	},
	welcomeHeader: {
		// margin: "auto",
		marginTop: "5vh",
		// background: "#fff",
		color: "#fff",
		padding: "1rem",
		// maxWidth: "600px",
		// marginRight: "10rem",
	},
	welcomeContent: {
		margin: "auto",
		// border: "1px solid black",
		// background: "#fff",
		padding: "1rem",
		// maxWidth: "900px",
		marginTop: "7rem",
		// boxShadow: "5px 5px #000",
	},
	sectionHeader: {
		paddingTop: "3rem",
		fontSize: "2.5rem",
	},
	joinBtn: {
		marginTop: "2rem",
		padding: "1.5rem",
		fontSize: "2rem",
		display: "flex",
		justifyContent: "center",
	},

	phoneImg: {
		// position: "absolute",
		width: "35rem",
		marginTop: "3rem",
		// background: "#fff",
		// marginLeft: "10rem",
		// right: 0,
	},

	aboutSection: {
		backgroundColor: "#000",
		color: "#fff",
		width: "100vw",
		minHeight: "80vh",
	},
	aboutContent: {
		maxWidth: "1200px",
		margin: "auto",
		display: "flex",
	},
	aboutImg: {
		background: `url("https://media.giphy.com/media/l0ErNdz1w5vt3YdZm/giphy.gif") no-repeat;`,
		height: "30rem",
		width: "33rem",
		margin: "1rem",
		backgroundSize: "auto",
	},

	videoSection: {
		backgroundColor: "#f4f4f4",
		width: "100vw",
		minHeight: "90vh",
	},
	videoContent: {
		maxWidth: "1200px",
		margin: "auto",
		display: "flex",
	},
	videoImg: {
		// background: `url("https://thehappypuppysite.com/wp-content/uploads/2018/11/pocket-beagle-long.jpg");`,
		height: "30rem",
		width: "33rem",
		margin: "1rem",
		backgroundSize: "contain",
	},
}
