import bg from "../imgs/bg1.svg"

// LANDING PAGE
export const global = {
	sticky: {
		position: "fixed",
		minWidth: "100vw",
	},
}
export const landingPageStyles = {
	welcomeHeader: {
		margin: "auto",
		marginTop: "5vh",
		background: "#fff",
		padding: "1rem",
		maxWidth: "600px",
	},
	welcomeContent: {
		margin: "auto",
		border: "1px solid black",
		background: "#fff",
		padding: "1rem",
		maxWidth: "900px",
		marginTop: "7rem",
		boxShadow: "5px 5px #000",
	},
	sectionHeader: {
		marginTop: "3rem",
		fontSize: "2.5rem",
	},
	joinBtn: {
		marginTop: "2rem",
		padding: "1.5rem",
		fontSize: "2rem",
	},
	landing: {
		minHeight: "100vh",
		backgroundColor: "#fff",
		width: "100vw",
		backgroundImage: `url(${bg})`,
		backgroundAttachment: "fixed",
		backgroundSize: "cover",
	},

	aboutSection: {
		backgroundColor: "#f4f4f4",
		width: "100vw",
		minHeight: "90vh",
	},
	aboutContent: {
		maxWidth: "1200px",
		margin: "auto",
		display: "flex",
	},
	aboutImg: {
		background: `url("https://techcrunch.com/wp-content/uploads/2015/06/gifgif.gif");`,
		height: "30rem",
		width: "33rem",
		margin: "1rem",
		backgroundSize: "contain",
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
