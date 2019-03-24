import Link from 'next/link'

const linkStyle = {
  marginRight: 15
}

export default function Header() {
  return (
    <div>
      <a style={linkStyle}>Bienvenue</a>
    </div>
  )
}
