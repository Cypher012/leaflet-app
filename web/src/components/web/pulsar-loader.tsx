import { Pulsar } from 'ldrs/react'
import 'ldrs/react/Pulsar.css'

const PulsarLoader = () => {
  return (
    <div className="flex justify-center items-center">
      <Pulsar size="80" speed="1.75" color="#317a4a" />
    </div>
  )
}

export default PulsarLoader