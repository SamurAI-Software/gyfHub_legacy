import React from "react"

// pages
import LandingPage from "./pages/landingPage"

// components
import Navbar from "./components/global/Navbar"

// Baseweb imports
import { Client as Styletron } from "styletron-engine-atomic"
import { Provider as StyletronProvider } from "styletron-react"
import { LightTheme, BaseProvider, styled } from "baseui"

const engine = new Styletron()
const Centered = styled("div", {
	display: "flex",
	overflowX: "hidden",
	flexDirection: "column",
	justifyContent: "center",
	alignItems: "center",
	height: "100%",
})

interface Props {}

const App: React.FC<Props> = () => {
	return (
		<React.Fragment>
			<StyletronProvider value={engine}>
				<BaseProvider theme={LightTheme}>
					{/* Nav */}
					<Navbar />
					<Centered>
						<LandingPage />
					</Centered>
				</BaseProvider>
			</StyletronProvider>
		</React.Fragment>
	)
}

export default App
