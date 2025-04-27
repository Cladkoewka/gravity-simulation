# Gravity Simulation

Gravity Simulation is a project that simulates gravitational interactions between planets, including orbits, velocities, and forces using Newtonian physics. The project uses the Raylib library for visualization on a 3D canvas, simulating the motion of planets and other bodies in the solar system.

![sim](https://github.com/user-attachments/assets/1fa03f51-6010-43c5-9076-9290ad0722fa)


## Description

This project models the solar system, including the Sun and planets, with the ability to control the time acceleration via a slider. The main physical principle is Newton's law of universal gravitation.

### Features:
- Visualization of planetary orbits in 3D.
- Control over time acceleration (from 1 day to 1 year per second).
- Simulation mode using real data for planet masses and orbital radii in the solar system.

## Project Structure

The project consists of several key parts:

1. **core**: The core module that contains the physical calculations and body models (planets, stars, moons, etc.).
2. **engine**: The simulation logic and time handling, including rendering using Raylib.
3. **graphics**: The module responsible for drawing objects on the screen, including planets and orbits.

