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
uniform float size = 100.0;
uniform float quality = 3.0;
uniform float directions = 32.0;
uniform float threshold = 0.2;

void main()
{
    float pi = 6.28318530718; // Pi*2
   
    vec2 radius = size/textureSize(texture0, 0);

    vec4 nearestColor = vec4(0);
    
    // Pixel color
    float alpha = texture(texture0, fragTexCoord).a;
    if (alpha == 1) {
        nearestColor = texture(texture0, fragTexCoord);
    }

    vec4 col = vec4(0);
    
    // Blur calculations
    for (float d=0.0; d<pi; d+=pi/directions) {
		for (float i=1.0/quality; i<=1.0; i+=1.0/quality) {
            col = texture(texture0, fragTexCoord+vec2(cos(d),sin(d))*radius*i);
			alpha += col.a;
            if (col.a == 1 && nearestColor.a == 0) {
                nearestColor = col;
            }
        }
    }
    
    // Output to screen
    alpha /= quality * directions - 15.0;
    if (alpha > threshold) {
        finalColor = nearestColor;
    } else {
        finalColor = vec4(0);
    }
}
