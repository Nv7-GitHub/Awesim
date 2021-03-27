#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;

// Output fragment color
out vec4 finalColor;

// NOTE: Add here your custom variables
uniform vec4 inpColor = vec4(0, 1, 1, 1);
uniform float size = 100.0;
uniform float quality = 3.0;
uniform float directions = 32.0;
uniform float threshold = 0.2;

// NOTE: Render size values must be passed from code
const float renderWidth = 800;
const float renderHeight = 450;

void main()
{
    float pi = 6.28318530718; // Pi*2
   
    vec2 Radius = size/textureSize(texture0, 0);
    
    // Pixel coloure
    vec4 color = texture(texture0, fragTexCoord);
    
    // Blur calculations
    for( float d=0.0; d<pi; d+=pi/directions)
    {
		for(float i=1.0/quality; i<=1.0; i+=1.0/quality)
        {
			color += texture( texture0, fragTexCoord+vec2(cos(d),sin(d))*Radius*i);		
        }
    }
    
    // Output to screen
    color /= quality * directions - 15.0;
    vec4 diff = vec4(1) - color;
    if ((diff.r + diff.g + diff.b)/3 > threshold) {
        color = inpColor;
    } else {
        color = vec4(1);
    }
    finalColor = color;
}
